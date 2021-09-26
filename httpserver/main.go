package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

const VERSION = "v1.0.0-alpha"

// header
func handleHeaderFunc(rw http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		rw.Header().Add(key, strings.Join(value, ""))
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("ok"))
}

// version
func handleVersionFunc(rw http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(version))
}

// healthz
func handleHealthzFunc(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("ok"))
}

// logger middleware
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Println("请求进来了")
		next.ServeHTTP(rw, r)
		// 状态码用原生的方式获取不到
		log.Printf("客户端IP：%s \n", r.RemoteAddr)
		log.Println("请求结束了")
	})
}

func init() {
	// 设置环境变量
	os.Setenv("VERSION", VERSION)
	// router
	router()
}

func router() {
	http.Handle("/header", loggerMiddleware(http.HandlerFunc(handleHeaderFunc)))
	http.Handle("/version", loggerMiddleware(http.HandlerFunc(handleVersionFunc)))
	http.Handle("/healthz", loggerMiddleware(http.HandlerFunc(handleHealthzFunc)))
}

func main() {
	log.Println("HTTPServer 启动中...")
	http.ListenAndServe(":8080", nil)
}
