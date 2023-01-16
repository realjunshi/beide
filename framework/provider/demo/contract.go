package demo

// Key Demo服务的key
const Key = "beide:demo"

// Service Demo 服务接口
type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
