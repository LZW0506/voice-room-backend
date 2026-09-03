// Package server 负责组装并创建 Gin HTTP 服务
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"voice-room-backend/internal/config"
	"voice-room-backend/internal/handler"
)

// New 创建带有 CORS 和业务路由的 Gin 引擎
func New(serviceConfig config.Config) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery(), corsMiddleware())

	httpHandler := handler.New(serviceConfig)
	engine.GET("/health", httpHandler.Health)
	engine.POST("/api/token", httpHandler.CreateToken)
	return engine
}

// corsMiddleware 设置浏览器客户端调用 Token API 所需的跨域响应头
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
		}
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusNoContent)
			c.Abort()
			return
		}
		c.Next()
	}
}
