FROM golang:1.14-alpine as builder

# 设置环境变量
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

# 设置工作区
WORKDIR /builder

# 代码
COPY cmd cmd
COPY internal internal
COPY go.mod go.mod

# 下载依赖包
RUN go mod download
# 编译
RUN CGO_ENABLE=0 GOOS=linux go build -o httpserver cmd/main.go

FROM alpine:latest as prod

# 设置工作区
WORKDIR /app

# 设置应用配置目录
ENV GS_SPRING_CONFIG_LOCATIONS configs

# 配置文件
COPY configs configs

# 复制已编译应用
COPY --from=builder /builder/httpserver httpserver

# 启动
ENTRYPOINT [ "/app/httpserver" ]
