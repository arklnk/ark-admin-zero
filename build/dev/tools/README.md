## 辅助工具的使用

开发模式

```
git clone https://github.com/arklnk/ark-admin-zero.git
```

```
cd ark-admin-zero
```

```
docker-compose -f docker-compose-dev.yml up -d
```

进入容器

```
docker exec -it api /bin/sh
```

映射工具

```
alias api="/server/build/dev/tools/api.sh"
alias model="/server/build/dev/tools/model.sh"
```

## api工具

```
api [api文件所在目录] [api文件]
```

例如api文件在：/app/core/cmd/api/core.api

如果要生成代码，则可以运行

```
api /app/core/cmd/api core.api
```

或者

```
cd /app/core/cmd/api
```

```
api ./ core.api
```

## model工具

```
model [表名] [model代码目录]
```

例如

```
model sys_user /app/core/model
```

