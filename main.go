package main

import (
	"beide/framework"
	"beide/framework/middleware"
	"net/http"
	"time"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Timeout(1 * time.Second))
	core.Use(middleware.Cost())
	//core.Use(middleware.Test1(), middleware.Test2())
	registerRouter(core)

	//subjectApi := core.Group("/user")
	//subjectApi.Use(middleware.Test3())

	server := &http.Server{
		Handler: core,
		Addr:    ":80",
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
