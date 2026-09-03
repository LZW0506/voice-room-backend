FROM golang:1.25-alpine AS builder

WORKDIR /src

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/voice-room-backend .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /out/voice-room-backend /app/voice-room-backend

EXPOSE 8787

ENTRYPOINT ["/app/voice-room-backend"]
