# 定制goctl名称、不知道改啥就直接写goctl
GO_CTL_NAME=goctl


# go-zero生成代码风格
#GO_ZERO_STYLE=goZero #推荐使用这一种风格
GO_ZERO_STYLE=gozero


GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
LDFLAGS := -s -w

.PHONY: test
test: # 运行项目测试
	go test -v --cover ./app/core/cmd/api/internal/..

.PHONY: fmt
fmt: # 格式化代码
	$(GOFMT) -w $(GOFILES)

.PHONY: gen-admin-api
gen-admin-api: # 生成 core-api 的代码
	$(GO_CTL_NAME) api go -api app/core/cmd/api/core.api -dir app/core/cmd/api  --style=$(GO_ZERO_STYLE)
	@echo "Generate core-api files successfully"

.PHONY: gen-all
gen-all: # 生成全部api和rpc代码
	make gen-admin-api
	#make gen-admin-rpc
	@echo "Generate all files successfully"

.PHONY: help
help: # 显示帮助
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
