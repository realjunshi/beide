package main

import (
	"beide/app/console"
	"beide/app/http"
	"beide/framework"
	"beide/framework/provider/app"
	"beide/framework/provider/config"
	"beide/framework/provider/distributed"
	"beide/framework/provider/env"
	"beide/framework/provider/kernel"
)

//func main() {
//	core := gin.New()
//
//	err := core.Bind(&app.BeideAppProvider{})
//	if err != nil {
//		return
//	}
//	err = core.Bind(&demo.DemoProvider{})
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	core.Use(gin.Recovery())
//	//core.Use(middleware.Timeout(10 * time.Second))
//	//core.Use(middleware.Cost())
//	//core.Use(middleware.Test1(), middleware.Test2())
//	beideHttp.Routes(core)
//
//	//subjectApi := core.Group("/user")
//	//subjectApi.Use(middleware.Test3())
//
//	server := &http.Server{
//		Handler: core,
//		Addr:    ":80",
//	}
//	// 启动服务的goroutine
//	go func() {
//		err := server.ListenAndServe()
//		if err != nil {
//			return
//		}
//	}()
//
//	// 当前的goroutine等待信号量
//	quit := make(chan os.Signal)
//	// 监控信号：SIGINT, SIGTERM, SIGQUIT
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
//	// 这里会阻塞当前goroutine等待信号
//	<-quit
//
//	// 调用Server.Shutdown graceful结束
//	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	if err := server.Shutdown(timeoutCtx); err != nil {
//		log.Fatal("Server Shutdown:", err)
//	}
//}

func main() {
	// 初始化服务容器
	container := framework.NewBeideContainer()
	// 绑定App服务提供者
	container.Bind(&app.BeideAppProvider{})
	// 后续初始化需要绑定的服务
	container.Bind(&env.HadeEnvProvider{})
	// 为什么无法走到register
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.BeideHadeKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
