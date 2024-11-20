package main

import (
	"fmt"
	"go-project/config"
	"go-project/router"
)

func main() {
	config.InitConfig()

	// 输出配置以便检查
	fmt.Println("Starting server on port:", config.AppConfig.App.Port)

	// 初始化路由
	r := router.SetupRouter()

	// 启动服务
	err := r.Run(config.AppConfig.App.Port)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
