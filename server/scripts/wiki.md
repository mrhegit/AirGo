

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

# 文档目录

<!-- TOC -->
* [1 面板部分功能展示](#1-面板部分功能展示)
* [2 安装教程](#2-安装教程)
  * [2-1 直接部署](#2-1-直接部署)
    * [2-1-1 安装`AirGo`核心](#2-1-1-安装-airgo-核心)
    * [2-1-2 配置文件](#2-1-2-配置文件)
    * [2-1-3 启动](#2-1-3-启动)
    * [2-1-4 配置ssl](#2-1-4-配置ssl)
    * [2-1-5 前端部署到Vercel，实现前后分离](#2-1-5-前端部署到-vercel-实现前后分离)
  * [2-2 使用宝塔面板部署](#2-2-使用宝塔面板部署)
    * [2-2-1 安装`AirGo`核心](#2-2-1-安装-airgo-核心)
    * [2-2-2 配置文件](#2-2-2-配置文件)
    * [2-2-3 启动](#2-2-3-启动)
    * [2-2-4 配置ssl](#2-2-4-配置ssl)
    * [2-2-5 前端部署到Vercel，实现前后分离](#2-2-5-前端部署到-vercel-实现前后分离)
* [3 对接XrayR](#3-对接xrayr)
* [4 其他说明](#4-其他说明)
<!-- TOC -->

<hr/>


# 1 面板部分功能展示

<table>
<tr>
    <td> <img src="">
    <td> <img src="">
</table>




# 2 安装教程

## 2-1 直接部署

### 2-1-1 安装`AirGo`核心

使用debian，ununtu，centos系统，执行以下命令，根据提示安装

```
bash <(curl -Ls https://raw.githubusercontent.com/ppoonk/AirGo/v2/server/scripts/install.sh)
```
### 2-1-2 配置文件

修改/usr/local/AirGo/config.yaml，根据自己的情况修改数据库、默认管理员等参数

### 2-1-3 启动
执行：`systemctl start AirGo`

浏览器访问：`http://ip:port`，即可打开网站

### 2-1-4 配置ssl

执行以下命令，根据提示配置ssl
```
https://
```
配置完成后，浏览器访问：`https://域名`，即可打开网站

如果需要自定义ssl，只需在安装目录（/usr/local/AirGo/）下，配置

### 2-1-5 前端部署到[Vercel](https://vercel.com)，实现前后分离



## 2-2 使用宝塔面板部署

### 2-2-1 安装`AirGo`核心

使用debian，ununtu，centos系统，执行以下命令，根据提示安装

```
https://
```
### 2-2-2 配置文件

修改/usr/local/AirGo/config.yaml，根据自己的情况修改数据库、默认管理员等参数，并且将http端口设置为非80端口，https设置为非443端口，避免和宝塔面板端口冲突

### 2-2-3 启动

### 2-2-4 配置ssl
先申请ssl，再开启反向代理，参考[宝塔网站开启反向代理时无法申请和自动续签SSL证书的解决办法](https://blog.csdn.net/qq_45576664/article/details/130171014)

### 2-2-5 前端部署到[Vercel](https://vercel.com)，实现前后分离







# 3 对接XrayR


# 4 其他说明
