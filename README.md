# ark-admin-zero

> Goland
>
> 版本：2021.2.2
>
> 链接：https://www.jetbrains.com/go/download/other.html
>
> golang
>
> 版本：1.18
>
> 链接：https://go.dev/dl
>
> go-zero
>
> 版本：1.4.0
>
> 链接：https://github.com/zeromicro/go-zero/releases/tag/v1.4.0
>
> goctl 
>
> 版本：1.4.0
>
> 链接：https://github.com/zeromicro/go-zero/releases/tag/tools/goctl/v1.4.0

### 搭建开发环境

```
docker-compose up -d
```

> 开发环境包含
>
> redis
>
> 密码：123456
>
> 
>
> mysql
>
> 账号：root
>
> 密码：root
>
> 
>
> phpmyadmin

### 导入数据

登录：http://127.0.0.1:8080

> 服务器：mysql
>
> 用户名：root
>
> 密码：root

新建数据库：ark_admin_zero，然后导入doc/sql/ark_admin_zero.sql

### 更新依赖

```shell
go mod tidy
```

### 安装热启动工具

```shell
go install github.com/cortesi/modd/cmd/modd
```

### 启动服务

```shell
modd
```

