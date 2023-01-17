package http

import (
	"beide/app/http/module/demo"
	"beide/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	err := demo.Register(r)
	if err != nil {
		return
	}
}
