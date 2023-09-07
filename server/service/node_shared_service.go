package service

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/utils/encrypt_plugin"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func ParseVMessLink(link string) *model.NodeShared {
	node := new(model.NodeShared)
	node.Enabled = true
	node.NodeType = "vmess"
	if strings.ToLower(link[:8]) == "vmess://" {
		link = link[8:]
	} else {
		return nil
	}
	if len(link) == 0 {
		return nil
	}
	jsonStr := encrypt_plugin.SubBase64Decode(link)
	if jsonStr == "" {
		return nil
	}
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		return nil
	}
	if version, ok := mapResult["v"]; ok {
		node.V = fmt.Sprintf("%v", version)
	}
	if ps, ok := mapResult["ps"]; ok {
		node.Remarks = fmt.Sprintf("%v", ps) //别名
	} else {
		return nil
	}
	if addr, ok := mapResult["add"]; ok {
		node.Address = fmt.Sprintf("%v", addr) //地址
	} else {
		return nil
	}
	if scy, ok := mapResult["scy"]; ok {
		node.Scy = fmt.Sprintf("%v", scy) //加密方式 auto,none,chacha20-poly1305,aes-128-gcm,zero
	} else {
		node.Scy = "auto"
	}
	if port, ok := mapResult["port"]; ok {
		value, err := strconv.ParseInt(fmt.Sprintf("%v", port), 10, 64)
		if err == nil {
			node.Port = value //端口
		} else {
			return nil
		}
	} else {
		return nil
	}

	if id, ok := mapResult["id"]; ok {
		node.UUID = fmt.Sprintf("%v", id) //uuid
	} else {
		return nil
	}
	if aid, ok := mapResult["aid"]; ok {
		if value, err := strconv.ParseInt(fmt.Sprintf("%v", aid), 10, 64); err == nil {
			node.Aid = value //额外id
		} else {
			return nil
		}
	} else {
		return nil
	}
	if net, ok := mapResult["net"]; ok {
		node.Network = fmt.Sprintf("%v", net) //传输协议
	} else {
		return nil
	}
	if type1, ok := mapResult["type"]; ok {
		node.Type = fmt.Sprintf("%v", type1)
	} else {
		return nil
	}

	//获取混淆
	if host, ok := mapResult["host"]; ok {
		node.Host = fmt.Sprintf("%v", host)
	} else {
		return nil
	}

	if path, ok := mapResult["path"]; ok {
		node.Path = fmt.Sprintf("%v", path)
	} else {
		return nil
	}
	if tls, ok := mapResult["tls"]; ok {
		node.Security = fmt.Sprintf("%v", tls)
	} else {
		return nil
	}
	if sni, ok := mapResult["sni"]; ok {
		node.Sni = fmt.Sprintf("%v", sni)
	}
	if alpn, ok := mapResult["alpn"]; ok {
		node.Alpn = fmt.Sprintf("%v", alpn)
	}
	return node
}

// vless://uuid@abc:80?encryption=none&type=ws&security=&host=www.10086.com&path=%2Fpath#%E5%B1%B1%E4%B8%9C
// [scheme:][//[userinfo@]host][/]path[?query][#fragment]
func ParseVLessLink(link string) *model.NodeShared {
	u, err := url.Parse(link)
	if err != nil {
		return nil
	}
	if u.User == nil || u.Scheme != "vless" {
		return nil
	}
	node := new(model.NodeShared)
	node.Enabled = true
	node.NodeType = "vless"

	//remarks
	node.Remarks = u.Fragment
	if node.Remarks == "" {
		node.Remarks = u.Host
	}
	//address
	node.Address = u.Hostname()
	//port
	node.Port, err = strconv.ParseInt(u.Port(), 10, 64)
	if err != nil {
		return nil
	}
	//uuid
	node.UUID = u.User.Username()

	//解析参数
	urlQuery := u.Query()
	if urlQuery.Get("flow") != "" {
		node.VlessFlow = urlQuery.Get("flow")
	}
	if urlQuery.Get("encryption") != "" {
		node.VlessEncryption = urlQuery.Get("encryption")
	}
	if urlQuery.Get("type") != "" {
		node.Network = urlQuery.Get("type")
	}
	if urlQuery.Get("security") != "" {
		node.Security = urlQuery.Get("security")
	}
	//获取混淆
	if urlQuery.Get("host") != "" {
		node.Host = urlQuery.Get("host")
	} else {
		return nil
	}

	if urlQuery.Get("path") != "" {
		node.Path = urlQuery.Get("path")
	}

	if urlQuery.Get("sni") != "" {
		node.Sni = urlQuery.Get("sni")
	}
	if urlQuery.Get("alpn") != "" {
		node.Alpn = urlQuery.Get("alpn")
	}
	if urlQuery.Get("allowInsecure") != "" {
		node.AllowInsecure = true
	}
	return node
}

func ParseTrojanLink(link string) *model.NodeShared {
	u, err := url.Parse(link)
	if err != nil {
		return nil
	}
	if u.User == nil || u.Scheme != "trojan" {
		return nil
	}
	node := new(model.NodeShared)
	node.Enabled = true
	node.NodeType = "trojan"
	//remarks
	node.Remarks = u.Fragment
	if node.Remarks == "" {
		node.Remarks = u.Host
	}
	//address
	node.Address = u.Hostname()
	//port
	node.Port, err = strconv.ParseInt(u.Port(), 10, 64)
	if err != nil {
		return nil
	}
	//uuid
	node.UUID = u.User.Username()

	//解析参数
	urlQuery := u.Query()
	if urlQuery.Get("network") != "" {
		node.Network = urlQuery.Get("network")
	}
	if urlQuery.Get("type") != "" {
		node.Type = urlQuery.Get("type")
	}
	//获取混淆
	if urlQuery.Get("host") != "" {
		node.Host = urlQuery.Get("host")
	} else {
		return nil
	}
	if urlQuery.Get("path") != "" {
		node.Path = urlQuery.Get("path")
	}
	if urlQuery.Get("tls") != "" {
		node.Security = urlQuery.Get("tls")
	}
	if urlQuery.Get("sni") != "" {
		node.Sni = urlQuery.Get("sni")
	}
	if urlQuery.Get("alpn") != "" {
		node.Alpn = urlQuery.Get("alpn")
	}
	if urlQuery.Get("allowInsecure") != "" {
		node.AllowInsecure = true
	}

	return node
}

// 新增节点
func NewNodeShared(n *[]model.NodeShared) error {
	return global.DB.Create(&n).Error
}

// 获取节点列表
func GetNodeSharedList() (*[]model.Node, error) {
	var list []model.Node
	err := global.DB.Model(&model.NodeShared{}).Find(&list).Error
	return &list, err

}

// 删除节点
func DeleteNodeShared(n *model.NodeShared) error {
	return global.DB.Delete(&n, n.ID).Error

}

func ParseUrl(url string) *[]model.NodeShared {
	//去掉前后空格
	url = strings.Trim(url, " \n")

	var data string
	if strings.HasPrefix(url, "v") || strings.HasPrefix(url, "t") {
		data = url
	} else {
		data = SubBase64Decode(url)
	}

	//s := strings.ReplaceAll(data, "\r\n", "\n")
	//s = strings.ReplaceAll(s, "\r", "\n")

	s := strings.ReplaceAll(data, "\r\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "vless", "\nvless")
	s = strings.ReplaceAll(s, "vmess", "\nvmess")
	s = strings.ReplaceAll(s, "trojan", "\ntrojan")

	//fmt.Println("s:", s)
	list := strings.Split(strings.TrimRight(s, "\n"), "\n")
	//fmt.Println("list:", list)

	var Nodes []model.NodeShared
	for _, v := range list {
		data := ParseLink(v)
		if data == nil {
			continue
		}
		Nodes = append(Nodes, *data)
	}
	return &Nodes
}

// 解析一条节点,vmess vless trojan
func ParseLink(link string) *model.NodeShared {
	//fmt.Println("解析一条链接", link)
	u, err := url.Parse(link)
	if err != nil {
		return nil
	}
	switch u.Scheme {
	case "vmess":
		if obj := ParseVMessLink(link); obj != nil {
			return obj
		}
	case "vless":
		if obj := ParseVLessLink(link); obj != nil {
			return obj
		}
	case "trojan":
		if obj := ParseTrojanLink(link); obj != nil {
			return obj
		}
	}
	return nil
}

// 对节点base64格式进行解析
func SubBase64Decode(str string) string {
	i := len(str) % 4
	switch i {
	case 1:
		str = str[:len(str)-1]
	case 2:
		str += "=="
	case 3:
		str += "="
	}
	//str = strings.Split(str, "//")[1]
	var data []byte
	var err error
	if strings.Contains(str, "-") || strings.Contains(str, "_") {
		data, err = base64.URLEncoding.DecodeString(str)

	} else {
		data, err = base64.StdEncoding.DecodeString(str)
		//data, err = base64.RawURLEncoding.DecodeString(str)
	}
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}
