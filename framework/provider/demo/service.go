package demo

import (
	"beide/framework"
	"fmt"
)

// DemoService 具体的接口实例
type DemoService struct {
	// 实现接口
	Service
	// 参数
	c framework.Container
}

// NewDemoService 实例化sevice
func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)

	fmt.Println("new demo service")

	return &DemoService{
		c: c,
	}, nil
}

// GetFoo 实现接口
func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "I am foo",
	}
}
