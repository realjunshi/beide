package main

import (
	"beide/framework/gin"
)

func UserLoginController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectListController")
}
