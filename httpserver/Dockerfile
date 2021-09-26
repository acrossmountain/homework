FROM golang:1.14-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

WORKDIR /app/web
COPY . .

RUN go mod download
RUN CGO_ENABLE=0 GOOS=linux go build -o httpserver

FROM alpine:latest
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/web/httpserver .
ENTRYPOINT [ "./httpserver" ]