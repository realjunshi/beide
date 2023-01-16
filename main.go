package main

import (
	"beide/framework/gin"
	"beide/framework/provider/demo"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()

	err := core.Bind(&demo.DemoServiceProvider{})
	if err != nil {
		log.Println(err)
		return
	}
	core.Use(gin.Recovery())
	//core.Use(middleware.Timeout(10 * time.Second))
	//core.Use(middleware.Cost())
	//core.Use(middleware.Test1(), middleware.Test2())
	registerRouter(core)

	//subjectApi := core.Group("/user")
	//subjectApi.Use(middleware.Test3())

	server := &http.Server{
		Handler: core,
		Addr:    ":80",
	}
	// 启动服务的goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
