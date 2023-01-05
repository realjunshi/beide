package main

import (
	"beide/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":80",
	}
	server.ListenAndServe()
}
