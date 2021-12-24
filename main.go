package main

import (
	"52lu/go-import-template/core"
	"52lu/go-import-template/initialize"
)

func main() {
	// 初始化配置
	initialize.InitViperConfig()
	// 启动服务
	core.RunServer()
}

