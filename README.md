# 声屿后端

这是基于 Go + Gin 的独立无状态 Token 服务与 LiveKit 单机运行项目，不使用数据库和登录系统

## 首次配置

```bash
cp .env.example .env
docker compose up -d
```

## 启动 Token 服务

无需提前构建，直接运行：

```bash
go run main.go
```

Token API 为 `POST /api/token`，健康检查为 `GET /health`

`LIVEKIT_URL` 用于后端容器访问同一 Compose 网络中的 LiveKit，默认是 `ws://livekit:7880`。`LIVEKIT_PUBLIC_URL` 是返回给 Tauri 客户端的地址，生产环境配置为 `ws://82.157.174.249:7880`

Tauri 客户端直接访问 `http://82.157.174.249:8787` 和 `ws://82.157.174.249:7880`

启动顺序如下：先在本目录执行 `docker compose up -d` 启动 Gin 与 LiveKit，再启动 Tauri 客户端

## 多节点扩展

当前 `livekit.yaml` 是单机路由配置。扩展多个 LiveKit 节点时增加 Redis 配置，并将各节点放在负载均衡之后；Token API 不保存状态，可以单独部署多个实例
