// Package main 是声屿 Gin Token 服务的启动入口
package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"voice-room-backend/internal/config"
	"voice-room-backend/internal/server"
)

// main 读取配置并启动 HTTP 服务
func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("未读取 .env，将使用当前环境变量：%v", err)
	}

	serviceConfig, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	engine := server.New(serviceConfig)
	address := fmt.Sprintf("0.0.0.0:%d", serviceConfig.Port)
	log.Printf("Token 服务已启动：http://localhost:%d", serviceConfig.Port)
	if err := engine.Run(address); err != nil {
		log.Fatal(err)
	}
}
