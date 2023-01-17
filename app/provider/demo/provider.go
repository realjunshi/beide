package demo

import (
	"beide/framework"
)

type DemoProvider struct {
	framework.ServiceProvider

	c framework.Container
}

func (sp *DemoProvider) Name() string {
	return Key
}

func (sp *DemoProvider) Register(c framework.Container) framework.NewInstance {
	return NewService
}

func (sp *DemoProvider) IsDefer() bool {
	return false
}

func (sp *DemoProvider) Params(c framework.Container) []interface{} {
	return []interface{}{sp.c}
}

func (sp *DemoProvider) Boot(c framework.Container) error {
	sp.c = c
	return nil
}
