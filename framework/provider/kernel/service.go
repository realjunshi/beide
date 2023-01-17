package kernel

import (
	"beide/framework/gin"
	"net/http"
)

// 引擎服务
type BeideKernelService struct {
	engine *gin.Engine
}

// NewBeideKernelService 初始化web引擎服务实例
func NewBeideKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &BeideKernelService{engine: httpEngine}, nil
}

// HttpEngine 返回web引擎
func (s *BeideKernelService) HttpEngine() http.Handler {
	return s.engine
}
