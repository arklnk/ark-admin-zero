![](imgs/README/ark-admin.png)

## 官方文档

传送门：[https://docs.arklnk.com](https://docs.arklnk.com)

## 在线体验

在线体验地址: [http://106.52.40.152:8080](http://106.52.40.152:8080)

| 账号 | 密码   | 备注     |
| ---- | ------ | -------- |
| demo | 123456 | 演示账号 |

- 如需体验更完整的项目功能可自行下载项目并运行体验或使用 docker-compose

## docker-compose

```sh
git clone https://github.com/arklnk/ark-admin-zero.git
```

```sh
cd ark-admin-zero
```

```
docker-compose up -d
```

登录地址：[http://127.0.0.1](http://127.0.0.1)

| 账号   | 密码   | 备注       |
| ------ | ------ | ---------- |
| arklnk | 123456 | 超级管理员 |
| demo   | 123456 | 演示账号   |

![login](https://docs.arklnk.com/images/zero/login.png)

![](https://docs.arklnk.com/images/zero/menu.png)

## 开发模式

### 环境要求

- golang   （ 版本：1.18）链接：[https://go.dev/dl](https://go.dev/dl)
- go-zero （版本：1.4.0） 链接：[https://github.com/zeromicro/go-zero/releases/tag/v1.4.0](https://github.com/zeromicro/go-zero/releases/tag/v1.4.0)
- goctl      （ 版本：1.4.0）链接：[https://github.com/zeromicro/go-zero/releases/tag/tools/goctl/v1.4.0](https://github.com/zeromicro/go-zero/releases/tag/tools/goctl/v1.4.0)

### clone项目

```sh
git clone https://github.com/arklnk/ark-admin-zero.git
```

```sh
cd ark-admin-zero
```

### 下载依赖

```sh
go mod tidy
```

### 热启动

```sh
go get github.com/cortesi/modd/cmd/modd
```

编辑热启动配置（项目根目录下的modd.conf）

window环境下

```conf
#core
app/core/**/*.* {
    prep: go build -o data/service/core-api.exe -v app/core/cmd/api/core.go
    daemon: data/service/core-api.exe -f app/core/cmd/api/etc/core-api.yaml
}
```

mac、linux环境下

```
#core
app/core/**/*.* {
    prep: go build -o data/service/core-api -v app/core/cmd/api/core.go
    daemon: data/service/core-api -f app/core/cmd/api/etc/core-api.yaml
}
```

> 注：modd开源地址https://github.com/cortesi/modd

运行（开发环境需要用到 redis 和 mysql，所以请先执行 docker-compose up -d 来启动容器）

```
modd
```

如输出以下信息则启动成功

```
15:17:03: prep: go build -o data/service/core-api.exe -v app/core/cmd/api/core.go
>> done (2.2430897s)
15:17:06: daemon: data/service/core-api.exe -f app/core/cmd/api/etc/core-api.yaml
>> starting...
Starting server at 0.0.0.0:7001...
```

# 欢迎 Star && PR

如果项目有帮助到你可以点个 Star 支持下。有更好的实现欢迎 PR。

## 浏览器支持

Modern browsers and Internet Explorer 10+.

| [![IE / Edge](https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png)](https://godban.github.io/browsers-support-badges/) IE / Edge | [![Firefox](https://raw.githubusercontent.com/alrra/browser-logos/master/src/firefox/firefox_48x48.png)](https://godban.github.io/browsers-support-badges/) Firefox | [![Chrome](https://raw.githubusercontent.com/alrra/browser-logos/master/src/chrome/chrome_48x48.png)](https://godban.github.io/browsers-support-badges/) Chrome | [![Safari](https://raw.githubusercontent.com/alrra/browser-logos/master/src/safari/safari_48x48.png)](https://godban.github.io/browsers-support-badges/) Safari |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| IE10, IE11, Edge                                             | last 2 versions                                              | last 2 versions                                              | last 2 versions                                              |

## 赞赏或交流

![交流](https://docs.arklnk.com/images/zero/contact.jpg)

## 致谢

[https://go-zero.dev/cn](https://go-zero.dev/cn)

[https://github.com/Mikaelemmmm/go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook)

[https://github.com/jinzhu/copier](https://github.com/jinzhu/copier)

[https://github.com/go-playground/validator](https://github.com/go-playground/validator)

[https://github.com/fangpenlin/avataaars-generator](https://github.com/fangpenlin/avataaars-generator)

[https://www.jetbrains.com/go](https://www.jetbrains.com/go)

