

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

使用debian，ununtu，centos系统，执行以下命令，根据提示申请证书

```
bash <(curl -Ls https://raw.githubusercontent.com/ppoonk/AirGo/v2/server/scripts/install.sh)
```

如果使用自定义ssl，只需在安装目录（/usr/local/AirGo/）下，配置 `air.cer`，`air.key`

### 2-1-5 前端部署到[Vercel](https://vercel.com)，实现前后分离

- fork本项目，修改`项目/web/.env`的`VITE_API_URL`字段为自己的后端地址
- 登录[Vercel](https://vercel.com)，Add New Project，参考下图配置
![image](https://telegraph-image.pages.dev/file/45b42a633b066deb52abb.png)
- 部署成功后，自定义域名即可（域名解析到76.76.21.21)

## 2-2 使用[宝塔面板](https://www.bt.cn/)部署

### 2-2-1 安装`AirGo`核心

同2-1-1

### 2-2-2 配置文件

修改/usr/local/AirGo/config.yaml，根据自己的情况修改数据库、默认管理员等参数，并且将**http端口设置为非80端口**，**https设置为非443端口**，避免和宝塔面板端口冲突

### 2-2-3 启动
同2-1-3

### 2-2-4 配置ssl
- 使用本项目提供的脚本申请，或者自定义ssl，同2-1-4
- 使用[宝塔面板](https://www.bt.cn/)申请，先申请ssl，再开启反向代理，参考[宝塔网站开启反向代理时无法申请和自动续签SSL证书的解决办法](https://blog.csdn.net/qq_45576664/article/details/130171014)

### 2-2-5 前端部署到[Vercel](https://vercel.com)，实现前后分离
同2-1-5

# 3 对接XrayR等后端
本项目使用sspanel同名api，所以对接只需将面板类型设置为`sspanel`即可

# 4 其他说明

