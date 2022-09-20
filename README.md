![](https://docs.arklnk.com/images/ark-admin.png)

English | [简体中文](README-cn.md)

## The official documentation

Doc: [https://docs.arklnk.com](https://docs.arklnk.com)

The back-end: [https://github.com/arklnk/ark-admin-vuenext](https://github.com/arklnk/ark-admin-vuenext)

The front end：[https://github.com/arklnk/ark-admin-zero](https://github.com/arklnk/ark-admin-zero)

## Online experience

demo: [http://arkadmin.si-yee.com](http://arkadmin.si-yee.com)

| Account | Password | Remark |
| ------- | -------- | ------ |
| demo    | 123456   | demo   |

- For more complete project functionality, download the project yourself and run the experience or use **docker-compose**

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

login：[http://127.0.0.1](http://127.0.0.1)

| Account | Password | Remark |
| ------- | -------- | ------ |
| arklnk  | 123456   | root   |
| demo    | 123456   | demo   |

![login](https://docs.arklnk.com/images/zero/login.png)

![](https://docs.arklnk.com/images/zero/menu.png)

## Development mode

### Environmental requirements

- golang   （ Version:1.18）link：[https://go.dev/dl](https://go.dev/dl)
- go-zero （Version:1.4.1） link：[https://github.com/zeromicro/go-zero/releases/tag/v1.4.1](https://github.com/zeromicro/go-zero/releases/tag/v1.4.1)
- goctl      （ Version:1.4.1）link：[https://github.com/zeromicro/go-zero/releases/tag/tools/goctl/v1.4.1](https://github.com/zeromicro/go-zero/releases/tag/tools/goctl/v1.4.1)

### Clone project

```sh
git clone https://github.com/arklnk/ark-admin-zero.git
```

```sh
cd ark-admin-zero
```

### Download the dependent

```sh
go mod tidy
```

### modd

```sh
go get github.com/cortesi/modd/cmd/modd
```

Edit the hot start configuration（modd.conf in the project root directory）

Windows

```conf
#core
app/core/**/*.* {
    prep: go build -o data/service/core-api.exe -v app/core/cmd/api/core.go
    daemon: data/service/core-api.exe -f app/core/cmd/api/etc/core-api.yaml
}
```

MAC or Linux 

```
#core
app/core/**/*.* {
    prep: go build -o data/service/core-api -v app/core/cmd/api/core.go
    daemon: data/service/core-api -f app/core/cmd/api/etc/core-api.yaml
}
```

run（The development environment will use Redis and mysql, so run **docker-compose up -d** to start the container）

```
modd
```

If the following information is displayed, the system is successfully started

```
15:17:03: prep: go build -o data/service/core-api.exe -v app/core/cmd/api/core.go
>> done (2.2430897s)
15:17:06: daemon: data/service/core-api.exe -f app/core/cmd/api/etc/core-api.yaml
>> starting...
Starting server at 0.0.0.0:7001...
```

# Star && PR

If the project is helpful you can click on the Star support. There are better implementations of welcome PR.

## Browser Support

Modern browsers and Internet Explorer 10+.

| [![IE / Edge](https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png)](https://godban.github.io/browsers-support-badges/) IE / Edge | [![Firefox](https://raw.githubusercontent.com/alrra/browser-logos/master/src/firefox/firefox_48x48.png)](https://godban.github.io/browsers-support-badges/) Firefox | [![Chrome](https://raw.githubusercontent.com/alrra/browser-logos/master/src/chrome/chrome_48x48.png)](https://godban.github.io/browsers-support-badges/) Chrome | [![Safari](https://raw.githubusercontent.com/alrra/browser-logos/master/src/safari/safari_48x48.png)](https://godban.github.io/browsers-support-badges/) Safari |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| IE10, IE11, Edge                                             | last 2 versions                                              | last 2 versions                                              | last 2 versions                                              |

## Thanks

[https://go-zero.dev/cn](https://go-zero.dev/cn)

[https://github.com/Mikaelemmmm/go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook)

[https://github.com/jinzhu/copier](https://github.com/jinzhu/copier)

[https://github.com/go-playground/validator](https://github.com/go-playground/validator)

[https://github.com/fangpenlin/avataaars-generator](https://github.com/fangpenlin/avataaars-generator)

[https://www.jetbrains.com/go](https://www.jetbrains.com/go)

[https://github.com/cortesi/modd](https://github.com/cortesi/modd)

