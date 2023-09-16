

<img width="200px" src="https://telegraph-image.pages.dev/file/c48a2f45ebf102dd66131.png" align="left"/>

# AirGo 前后端分离机场面板，简单易用

前端：vue-next-admin框架（vue，typescript，pinia，vite）

后端：golang，gin，gorm

![License](https://img.shields.io/badge/License-GPL_v3.0-red)
![Go](https://img.shields.io/badge/Golang-orange?logo=Go&logoColor=white)
![Gorm](https://img.shields.io/badge/Gorm-yellow&logo=gorm)
![Gin](https://img.shields.io/badge/Gin-green?logo=)
![Vue](https://img.shields.io/badge/Vue.js-00b6ff?logo=vuedotjs&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-blue?logo=TypeScript&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-purple?logo=linux&logoColor=white)

<hr/>

# 1 面板部分功能展示

<table>
<tr>
    <td> <img src="https://telegraph-image.pages.dev/file/0a377278d4d264a8c0521.png">
    <td> <img src="https://telegraph-image.pages.dev/file/533bff998724b7bd87ef0.png">
<tr>
    <td> <img src="https://telegraph-image.pages.dev/file/1a8eb3c9bf615ea4c4cd0.png">
    <td> <img src="https://telegraph-image.pages.dev/file/b17bf699f6cc3e47b0d1f.png">
</table>




# 2 安装教程
>前后分离开发，并且静态资源嵌入可执行文件中，所以部署非常灵活。下面仅列举常用的几种方式，更多部署方式请自行实践

## 2-1 直接部署

### 2-1-1 安装`AirGo`核心

使用debian，ununtu，centos系统，执行以下命令，根据提示安装AirGo核心

```
bash <(curl -Ls https://raw.githubusercontent.com/ppoonk/AirGo/v2/server/scripts/install.sh)
```
### 2-1-2 配置文件

修改/usr/local/AirGo/config.yaml，根据自己的情况修改数据库、默认管理员等参数

### 2-1-3 启动
- 执行：`systemctl start AirGo`，浏览器访问：`http://ip:port`

- 或者终端输入：`AirGo`，根据提示启动

### 2-1-4 配置ssl

- 使用debian，ununtu，centos系统，执行以下命令，根据提示申请证书

```
bash <(curl -Ls https://raw.githubusercontent.com/ppoonk/AirGo/v2/server/scripts/install.sh)
```

- 如果使用自定义ssl，只需在安装目录（/usr/local/AirGo/）下，配置 `air.cer`，`air.key`
- 配置完ssl，需重启服务生效

### 2-1-5 前端部署到[Vercel](https://vercel.com)，实现前后分离

- fork本项目，修改`项目/web/.env`的`VITE_API_URL`字段为自己的后端地址
- 登录[Vercel](https://vercel.com)，Add New Project，参考下图配置
![image](https://telegraph-image.pages.dev/file/45b42a633b066deb52abb.png)
- 部署成功后，自定义域名即可（域名解析到76.76.21.21)

## 2-2 使用[宝塔面板](https://www.bt.cn/)部署

### 2-2-1 安装`AirGo`核心

同2-1-1

### 2-2-2 配置文件

- 修改/usr/local/AirGo/config.yaml，根据自己的情况修改数据库、默认管理员等参数，并且将**http端口设置为非80端口**，**https设置为非443端口**，避免和宝塔面板端口冲突

### 2-2-3 启动
同2-1-3

### 2-2-4 配置ssl
- 使用本项目提供的脚本申请，或者自定义ssl，同2-1-4
- 使用[宝塔面板](https://www.bt.cn/)申请，先申请ssl，再开启反向代理，参考[宝塔网站开启反向代理时无法申请和自动续签SSL证书的解决办法](https://blog.csdn.net/qq_45576664/article/details/130171014)

### 2-2-5 前端部署到[Vercel](https://vercel.com)，实现前后分离
同2-1-5

## 2-4 Docker部署

- 在合适的目录新建配置文件，例如：/$PWD/air/config.yaml，配置文件内容如下：

```
system:
  admin-email: admin@oicq.com
  admin-password: adminadmin
  http-port: 80
  https-port: 443
  db-type: sqlite
mysql:
  address: mysql.sql.com
  port: 3306
  config: charset=utf8mb4&parseTime=True&loc=Local
  db-name: imdemo
  username: imdemo
  password: xxxxxx
  max-idle-conns: 10
  max-open-conns: 100
sqlite:
  path: ./air.db

```
- 根据自己的需求，修改配置文件，启动docker命令参考如下：

```
docker run -tid \
  -v $PWD/air/config.yaml:/air/config.yaml \
  -p 80:80 \
  -p 443:443 \
  --name airgo \
  --restart always \
  --privileged=true \
  ppoiuty/airgo:latest
```



## 2-5 手动安装
linux,darwin 下载对应压缩包，解压后启动：`./AirGo -start`

# 5 关于套餐监控
>本人是个免流佬，这个功能纯粹是免流用的，可能大部分人用不到，可以跳过。免流一般用青龙面板或者1ts这个app监控流量，我参考部分开源项目，然后用安卓的KWGT这个组件app，写了个简单的流量监控组件。
支持联通，电信手机号。
## 5-1 功能简介
<table>
<tr>
    <td> <img src="https://wiki.airgoo.link/assets/home.368f2b00.jpeg">
</table>

## 5-1 下载

[KWGT下载](https://www.123pan.com/s/oIT9-qhyxH.html)

[AirGo插件下载---<font color=red>更新时间：2023-08-31</font>](https://www.123pan.com/s/oIT9-QxyxH.html)


## 5-2 安装

- 安装KWGT，在桌面添加Kustom Widget，先选择4x4尺寸，在添加插件后可自行调整大小
![kwgt1](https://wiki.airgoo.link/assets/kwgt1.efa3c135.jpg)
- 注册登录机场。点击菜单栏-流量监控，登录自己的手机号，然后复制url
- 打开KWGT，点击 `导入`，导入AirGo插件-`AirGo.kwgt`
  ![kwgt1](https://wiki.airgoo.link/assets/kwgt2.8ecfb18e.jpg)
- 点击-编辑
  ![kwgt1](https://wiki.airgoo.link/assets/kwgt3.6465d5e9.jpg)

- 点击-全局变量-url
![kwgt1](https://wiki.airgoo.link/assets/kwgt4.a953bb41.jpg)

- 在公式编辑器里粘贴上面复制的url，<font color=red>点击右上角对号保存！然后再点击右上角保存按钮再次保存！</font>
![kwgt1](https://wiki.airgoo.link/assets/kwgt5.0ef0f989.jpg)

- 然后回到桌面，点击一次手动刷新按钮，等待一会，看是否刷新出流量使用详情
![kwgt1](https://wiki.airgoo.link/assets/kwgt6.7813848b.jpg)



# 4 对接XrayR等后端
本项目使用sspanel同名api，所以对接只需将面板类型设置为`sspanel`即可

# 5 开发注意事项
- 手动编译，脚本在`项目/server/scripts/install.sh`

# 6 其他说明



