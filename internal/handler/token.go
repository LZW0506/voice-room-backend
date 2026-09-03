// Package handler 提供 HTTP 接口处理逻辑
package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/livekit/protocol/auth"

	"voice-room-backend/internal/config"
)

// TokenRequest 是创建 LiveKit Token 的请求体
type TokenRequest struct {
	// Room 是要加入的房间名称
	Room string `json:"room"`
	// Identity 是基于机器码摘要生成的参与者身份
	Identity string `json:"identity"`
	// Name 是客户端显示昵称
	Name string `json:"name"`
}

// TokenResponse 是创建 LiveKit Token 的响应体
type TokenResponse struct {
	// Token 是 LiveKit 访问令牌
	Token string `json:"token"`
	// URL 是 LiveKit WebSocket 地址
	URL string `json:"url"`
}

// Handler 保存 HTTP 接口处理所需的服务配置
type Handler struct {
	// config 是服务运行配置
	config config.Config
}

// New 创建 HTTP 接口处理器
func New(serviceConfig config.Config) *Handler {
	return &Handler{config: serviceConfig}
}

// Health 返回服务健康状态
func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// CreateToken 校验请求并签发 LiveKit 访问令牌
func (h *Handler) CreateToken(c *gin.Context) {
	var request TokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求格式不正确"})
		return
	}

	room := readText(request.Room, 64)
	identity := readText(request.Identity, 64)
	name := readText(request.Name, 24)
	if room == "" || identity == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "房间名称、设备身份和昵称不能为空"})
		return
	}

	accessToken := auth.NewAccessToken(h.config.APIKey, h.config.APISecret).
		SetIdentity(identity).
		SetName(name).
		SetValidFor(2 * time.Hour)
	grant := &auth.VideoGrant{Room: room, RoomJoin: true}
	grant.SetCanPublish(true)
	grant.SetCanSubscribe(true)
	grant.SetCanPublishData(false)
	grant.SetCanUpdateOwnMetadata(true)
	accessToken.SetVideoGrant(grant)

	token, err := accessToken.ToJWT()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "创建房间通行证失败"})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{Token: token, URL: h.config.LiveKitPublicURL})
}

// readText 清理请求文本并限制其最大字符数
func readText(value string, maxLength int) string {
	value = strings.TrimSpace(value)
	runes := []rune(value)
	if len(runes) > maxLength {
		runes = runes[:maxLength]
	}
	return string(runes)
}
