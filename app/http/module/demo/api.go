package demo

import (
	"beide/app/provider/demo"
	"beide/framework/contract"
	"beide/framework/gin"
)

type DemoApi struct {
	service *Service
}

func (a DemoApi) Demo(c *gin.Context) {
	//appService := c.MustMake(contract.AppKey).(contract.App)
	//baseFolder := appService.BaseFolder()
	//
	//log.Println("baseFolder:", baseFolder)
	////users := api.service.GetUsers()
	////usersDTO := UserModelsToUserDTOs(users)
	////c.JSON(200, usersDTO)
	//c.JSON(200, baseFolder)

	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	//password := configService.IsExist(contract.ConfigKey)
	//log.Println(r)
	password := configService.GetString("database.mysql.password") // 打印出来

	logger := c.MustMakeLog()
	logger.Info(c, "demo test error", map[string]interface{}{
		"api":  "demo/demo",
		"user": "jianfengye",
	})

	c.JSON(200, password)
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	err := r.Bind(&demo.DemoProvider{})
	if err != nil {
		return err
	}

	r.GET("/demo/demo", api.Demo)
	//r.GET("/demo/demo2", api.Demo2)
	//r.POST("/demo/demo_post", api.DemoPost)
	return nil
}
