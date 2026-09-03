// Package config 提供服务运行配置的读取与校验能力
package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config 保存 Token 服务运行所需的配置
type Config struct {
	// Port 是 HTTP 服务监听端口
	Port int
	// LiveKitURL 是后端访问本机 LiveKit 的 WebSocket 地址
	LiveKitURL string
	// LiveKitPublicURL 是返回给局域网客户端的 LiveKit WebSocket 地址
	LiveKitPublicURL string
	// APIKey 是 LiveKit API Key
	APIKey string
	// APISecret 是 LiveKit API Secret
	APISecret string
}

// Load 从环境变量读取配置并校验 LiveKit 密钥
func Load() (Config, error) {
	port, err := readPort()
	if err != nil {
		return Config{}, err
	}

	config := Config{
		Port:             port,
		LiveKitURL:       readEnv("LIVEKIT_URL", "ws://localhost:7880"),
		LiveKitPublicURL: readEnv("LIVEKIT_PUBLIC_URL", "ws://82.157.174.249:7880"),
		APIKey:           os.Getenv("LIVEKIT_API_KEY"),
		APISecret:        os.Getenv("LIVEKIT_API_SECRET"),
	}
	if config.APIKey == "" || config.APISecret == "" {
		return Config{}, fmt.Errorf("缺少 LIVEKIT_API_KEY 或 LIVEKIT_API_SECRET")
	}
	return config, nil
}

// readPort 读取并解析 HTTP 服务端口
func readPort() (int, error) {
	portText := readEnv("PORT", "8787")
	port, err := strconv.Atoi(portText)
	if err != nil || port < 1 || port > 65535 {
		return 0, fmt.Errorf("PORT 必须是 1 到 65535 之间的数字")
	}
	return port, nil
}

// readEnv 读取环境变量，不存在时返回默认值
func readEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return defaultValue
	}
	return value
}
