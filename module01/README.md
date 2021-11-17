# 练习

- 编写一个 HTTP 服务器。
1. 接收客户端 request，并将 request 中带的 header 写入 response header 中。
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header。
3. Server 端记录访问日志包括客户端IP，HTTP 返回码，输出到 Server 端的标准输出。
4. 当访问 localhost/healthz 时，响应 200。