package service

import (
	"AirGo/global"
	"AirGo/model"
	"encoding/base64"
	"encoding/json"
	"gorm.io/gorm"
	"net/url"
	"strconv"
	"strings"

	//"gopkg.in/yaml.v2"
	"gopkg.in/yaml.v2"
)

// 节点信息
func SSNodeInfo(nodeID int64) (model.SSNodeInfo, error) {
	var node model.Node
	err := global.DB.Where("id = ? and enabled = true", nodeID).First(&node).Error //节点号 是否启用
	if err != nil {
		return model.SSNodeInfo{}, err
	}
	var nodeInfo model.SSNodeInfo
	//nodeInfo.NodeGroup = 0
	//nodeInfo.NodeClass = 0
	//nodeInfo.MuOnly = 1
	nodeInfo.NodeSpeedlimit = node.NodeSpeedlimit
	nodeInfo.TrafficRate = node.TrafficRate

	switch node.NodeType {
	case "vless":
		nodeInfo.Sort = 15
	case "vmess":
		nodeInfo.Sort = 11
	case "trojan":
		nodeInfo.Sort = 14

	}
	switch node.NodeType {
	case "vmess": //vmess
		if node.Type == "http" && node.Network == "tcp" {
			nodeInfo.Server = node.Address + ";" + strconv.FormatInt(node.Port, 10) + ";" + strconv.FormatInt(node.Aid, 10) + ";" + node.Network + ";" + node.Security + ";path=" + node.Path + "|host=" + node.Host + ";headertype=http"
		} else if node.Network == "grpc" && node.Security != "" {
			nodeInfo.Server = node.Address + ";" + strconv.FormatInt(node.Port, 10) + ";" + strconv.FormatInt(node.Aid, 10) + ";" + node.Network + ";" + node.Security + ";path=" + node.Path + "|host=" + node.Host + "|servicename=mygrpc"
		}
		nodeInfo.Server = node.Address + ";" + strconv.FormatInt(node.Port, 10) + ";" + strconv.FormatInt(node.Aid, 10) + ";" + node.Network + ";" + node.Security + ";path=" + node.Path + "|host=" + node.Host
	case "vless": //vless
		if node.Type == "http" && node.Network == "tcp" {
			nodeInfo.Server = node.Address + ";" + strconv.FormatInt(node.Port, 10) + ";" + strconv.FormatInt(node.Aid, 10) + ";" + node.Network + ";" + node.Server + ";path=" + node.Path + "|host=" + node.Host + ";headertype=http" + "|enable_vless=true"
		} else if node.Network == "grpc" && node.Security != "" {
			nodeInfo.Server = node.Address + ";" + strconv.FormatInt(node.Port, 10) + ";" + strconv.FormatInt(node.Aid, 10) + ";" + node.Network + ";" + node.Security + ";path=" + node.Path + "|host=" + node.Host + "|servicename=mygrpc" + "|enable_vless=true"
		}
		nodeInfo.Server = node.Address + ";" + strconv.FormatInt(node.Port, 10) + ";" + strconv.FormatInt(node.Aid, 10) + ";" + node.Network + ";" + node.Security + ";path=" + node.Path + "|host=" + node.Host + "|enable_vless=true"

	case "trojan": //trojan
		if node.Network == "grpc" {
			nodeInfo.Server = node.Address + ":" + strconv.FormatInt(node.Port, 10) + "|host=" + node.Host + "|grpc=1|servicename=mygrpc"
		}
		nodeInfo.Server = node.Address + ":" + strconv.FormatInt(node.Port, 10) + "|host=" + node.Host
	}
	return nodeInfo, nil

}

// 获取订阅
func GetUserSub(url string, subType string) string {
	//查找用户
	var u model.User
	err := global.DB.Where("subscribe_url = ? and sub_status = 1 and d + u < t", url).First(&u).Error
	if err != nil {
		return ""
	}
	//根据goodsID 查找具体的节点
	var goods model.Goods
	err = global.DB.Where("id = ?", u.SubscribeInfo.GoodsID).Preload("Nodes", func(db *gorm.DB) *gorm.DB { return db.Order("node_order") }).Find(&goods).Error
	// 计算剩余天数，流量
	//fmt.Println("根据goodsID 查找具体的节点", goods)
	expiredTime := u.SubscribeInfo.ExpiredAt.Format("2006-01-02")
	expiredBd1 := (float64(u.SubscribeInfo.T - u.SubscribeInfo.U - u.SubscribeInfo.D)) / 1024 / 1024 / 1024
	expiredBd2 := strconv.FormatFloat(expiredBd1, 'f', 2, 64)
	name := "到期时间:" + expiredTime + "  |  剩余流量:" + expiredBd2 + "GB"
	var firstSubNode = model.Node{
		Remarks:  name,
		Address:  global.Server.System.SubName,
		Port:     6666,
		Aid:      0,
		Network:  "ws",
		Enabled:  true,
		NodeType: "vmess",
	}
	//插入计算剩余天数，流量的第一条节点
	goods.Nodes = append(goods.Nodes, model.Node{})
	copy(goods.Nodes[1:], goods.Nodes[0:])
	goods.Nodes[0] = firstSubNode
	//再插入共享的节点
	nodeList, err := GetNodeSharedList()
	if err == nil {
		for _, v := range *nodeList {
			goods.Nodes = append(goods.Nodes, v)
		}
	}
	//fmt.Println("nodes:", goods.Nodes)
	//根据subType生成不同客户端订阅 1:v2rayng 2:clash 3 shadowrocket 4 Quantumult X
	switch subType {
	case "1":
		return V2rayNGSubscribe(&goods.Nodes, u.UUID.String(), u.SubscribeInfo.Host)
	case "2":
		return ClashSubscribe(&goods.Nodes, u.UUID.String(), u.SubscribeInfo.Host)
	case "3":
		//return ShadowRocketSubscribe(&goods.Nodes, u.UUID.String(), u.SubscribeInfo.Host, name)
		return V2rayNGSubscribe(&goods.Nodes, u.UUID.String(), u.SubscribeInfo.Host)
	case "4":
		//return QxSubscribe(&goods.Nodes, u.UUID.String(), u.SubscribeInfo.Host)
		return V2rayNGSubscribe(&goods.Nodes, u.UUID.String(), u.SubscribeInfo.Host)
	}
	return ""
}

// v2rayNG 订阅
func V2rayNGSubscribe(nodes *[]model.Node, uuid, host string) string {
	// 遍历，根据node sort 节点类型 生成订阅
	var subArr []string

	for _, v := range *nodes {
		//剔除禁用节点
		if !v.Enabled {
			continue
		}
		if host == "" {
			host = v.Host
		}
		switch v.NodeType {
		case "vmess":
			if res := V2rayNGVmess(v, uuid, host); res != "" {
				subArr = append(subArr, res)
			}
			continue
		case "vless":
			if res := V2rayNGVless(v, uuid, host); res != "" {
				subArr = append(subArr, res)
			}
			continue
		case "trojan":
			if res := V2rayNGTrojan(v, uuid, host); res != "" {
				subArr = append(subArr, res)
			}
			continue
		}
	}
	return base64.StdEncoding.EncodeToString([]byte(strings.Join(subArr, "\r\n")))
}

// clash 订阅
func ClashSubscribe(nodes *[]model.Node, uuid, host string) string {
	var proxiesArr []model.ClashProxy
	//所有节点名称数组
	var nameArr []string
	for _, v := range *nodes {
		//剔除禁用节点
		if !v.Enabled {
			continue
		}
		if host == "" {
			host = v.Host
		}
		//
		nameArr = append(nameArr, v.Remarks)

		switch v.NodeType {
		case "vmess":
			proxy := ClashVmess(v, uuid, host)
			proxiesArr = append(proxiesArr, proxy)
		case "vless":
			proxy := ClashVmess(v, uuid, host)
			proxiesArr = append(proxiesArr, proxy)
		case "trojan":
			proxy := ClashTrojan(v, uuid, host)
			proxiesArr = append(proxiesArr, proxy)
		}

	}
	var proxyGroup model.ClashProxyGroup
	proxyGroup.Name = global.Server.System.SubName
	proxyGroup.Type = "select"
	proxyGroup.Proxies = nameArr

	var clashYaml model.ClashYaml
	clashYaml.Port = 7890
	clashYaml.SocksPort = 7891
	clashYaml.RedirPort = 7892
	clashYaml.AllowLan = false
	clashYaml.Mode = "rule"
	clashYaml.LogLevel = "silent"
	clashYaml.ExternalController = "0.0.0.0:9090"
	clashYaml.Secret = ""
	clashYaml.Proxies = proxiesArr
	clashYaml.ProxyGroups = append(clashYaml.ProxyGroups, proxyGroup)
	clashYaml.Rules = append(clashYaml.Rules, "MATCH,"+global.Server.System.SubName)
	res, err := yaml.Marshal(clashYaml)
	if err != nil {
		global.Logrus.Error("yaml.Marshal error:", err)
		return ""
	}
	return string(res)

}

// ShadowRocket 订阅
func ShadowRocketSubscribe(nodes *[]model.Node, uuid, host, name string) string {
	// 遍历，根据node sort 节点类型 生成订阅
	var subArr []string

	for k, v := range *nodes {
		//剔除禁用节点
		if k == 0 || !v.Enabled {
			continue
		}
		if host == "" {
			host = v.Host
		}
		switch v.NodeType {
		case "vmess":
			if res := ShadowRocketVmess(v, uuid, host); res != "" {
				subArr = append(subArr, res)
			}
			continue
			// case 15 :
			// 	res:=GenerateVless()
		}
	}
	return "STATUS=" + name + "\r\n" + "REMARKS=" + global.Server.System.SubName + "\r\n" + strings.Join(subArr, "\r\n")

}

// Quantumult X 订阅
func QxSubscribe(nodes *[]model.Node, uuid, host string) string {
	var nodeArr []string
	for _, v := range *nodes {
		//剔除禁用节点
		if !v.Enabled {
			continue
		}
		if host == "" {
			host = v.Host
		}
		protocolType := ""
		switch v.NodeType {
		case "vmess":
			protocolType = "vmess="
		}
		str := protocolType + v.Address + ":" + strconv.FormatInt(v.Port, 10) + ", method=" + ", password=" + uuid + ", obfs=" + v.Network + ", obfs-uri=" + v.Path + ", obfs-host" + v.Host + ", tag=" + v.Remarks
		nodeArr = append(nodeArr, str)
	}
	return strings.Join(nodeArr, "\r\n")

}

// generate v2rayNG vmess
func V2rayNGVmess(node model.Node, uuid, host string) string {
	var vmess model.Vmess
	vmess.V = node.V
	vmess.Name = node.Remarks
	if node.EnableTransfer {
		vmess.Address = node.TransferAddress
		vmess.Port = strconv.FormatInt(node.TransferPort, 10)
	} else {
		vmess.Address = node.Address
		vmess.Port = strconv.FormatInt(node.Port, 10)
	}
	vmess.Uuid = uuid
	vmess.Aid = strconv.FormatInt(node.Aid, 10)
	vmess.Net = node.Network
	vmess.Disguisetype = node.Type
	vmess.Host = host
	vmess.Path = node.Path
	vmess.Tls = node.Security
	vmessMarshal, err := json.Marshal(vmess)
	if err != nil {
		return ""
	}
	vmessStr := base64.StdEncoding.EncodeToString([]byte(vmessMarshal))
	return "vmess://" + vmessStr
}

// generate  v2rayng vless
func V2rayNGVless(node model.Node, uuid, host string) string {
	path := url.QueryEscape(node.Path)
	name := url.QueryEscape(node.Remarks)
	var address, port string
	if node.EnableTransfer {
		address = node.TransferAddress
		port = strconv.FormatInt(node.TransferPort, 10)
	} else {
		address = node.Address
		port = strconv.FormatInt(node.Port, 10)
	}
	str := "vless://" + uuid + "@" + address + ":" + port + "?encryption=" + node.Scy + "&type=" + node.Network + "&security=" + node.Security + "&host=" + host + "&path=" + path
	if node.Security == "tls" || node.Security == "reality" {
		return str + "&sni=" + node.Sni + "#" + name
	}
	return str + "#" + name
}

// generate  v2rayng trojan
func V2rayNGTrojan(node model.Node, uuid, host string) string {
	//trojan://59405054-d6d2-47e1-8f99-b7296be5e7a1@114.114.114.114:80?allowInsecure=0#%E6%B5%8B%E8%AF%952
	path := url.QueryEscape(node.Path)
	name := url.QueryEscape(node.Remarks)
	var address, port string
	if node.EnableTransfer {
		address = node.TransferAddress
		port = strconv.FormatInt(node.TransferPort, 10)
	} else {
		address = node.Address
		port = strconv.FormatInt(node.Port, 10)
	}
	str := "trojan://" + uuid + "@" + address + ":" + port + "?security=" + node.Security + "&headerType=" + node.Type + "&type=" + node.Network + "&path=" + path + "&host=" + host
	if node.Security == "tls" || node.Security == "reality" {
		return str + "&sni=" + node.Sni + "#" + name
	}
	return str + "#" + name
}

// generate  Clash vmess vless
func ClashVmess(v model.Node, uuid, host string) model.ClashProxy {
	var proxy model.ClashProxy
	switch v.NodeType {
	case "vmess":
		proxy.Type = "vmess"
	case "vless":
		proxy.Type = "vless"
	}
	if v.EnableTransfer {
		proxy.Server = v.TransferAddress
		proxy.Port = strconv.FormatInt(v.TransferPort, 10)
	} else {
		proxy.Server = v.Address
		proxy.Port = strconv.FormatInt(v.Port, 10)
	}
	proxy.Name = v.Remarks
	proxy.Uuid = uuid
	proxy.Alterid = strconv.FormatInt(v.Aid, 10)
	proxy.Cipher = "auto"
	proxy.Udp = true
	proxy.Network = v.Network
	proxy.WsPath = v.Path
	proxy.WsHeaders.Host = host
	proxy.WsOpts.Path = v.Path
	proxy.WsOpts.Headers = make(map[string]string, 1)
	proxy.WsOpts.Headers["Host"] = host
	if v.Security != "" {
		proxy.Tls = true
		proxy.Sni = v.Sni
	}
	return proxy
}

// generate  Clash trojan
func ClashTrojan(v model.Node, uuid, host string) model.ClashProxy {
	var proxy model.ClashProxy
	if v.EnableTransfer {
		proxy.Server = v.TransferAddress
		proxy.Port = strconv.FormatInt(v.TransferPort, 10)
	} else {
		proxy.Server = v.Address
		proxy.Port = strconv.FormatInt(v.Port, 10)
	}
	proxy.Type = "trojan"
	proxy.Password = uuid
	proxy.Name = v.Remarks
	proxy.Uuid = uuid
	proxy.Alterid = strconv.FormatInt(v.Aid, 10)
	proxy.Cipher = "auto"
	proxy.Udp = true
	proxy.Network = v.Network
	proxy.WsPath = v.Path
	proxy.WsHeaders.Host = host
	proxy.WsOpts.Path = v.Path
	proxy.WsOpts.Headers = make(map[string]string, 1)
	proxy.WsOpts.Headers["Host"] = host
	if v.Security != "" {
		proxy.Tls = true
		proxy.Sni = v.Sni
	}
	return proxy
}

// generate ShadowRocket vmess
func ShadowRocketVmess(node model.Node, uuid, host string) string {
	//nameStr:="chacha20-poly1305:"
	var address, port string
	if node.EnableTransfer {
		address = node.TransferAddress
		port = strconv.FormatInt(node.TransferPort, 10)
	} else {
		address = node.Address
		port = strconv.FormatInt(node.Port, 10)
	}
	name := node.Scy + ":" + uuid + "@" + address + ":" + port
	nameStr := base64.StdEncoding.EncodeToString([]byte(name))
	netType := "websocket"
	switch node.Network {
	case "ws":
		netType = "websocket"
	}
	remarksStr := "?remarks=" + url.QueryEscape(node.Remarks) + "&obfsParam=" + host + "&path=" + node.Path + "&obfs=" + netType + "&alterId=" + strconv.FormatInt(node.Aid, 10)
	return "vmess://" + nameStr + remarksStr

}
