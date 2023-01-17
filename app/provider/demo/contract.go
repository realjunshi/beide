package demo

// Key Demo服务的key
const Key = "beide:demo"

type IService interface {
	GetAllStudent() []Student
}

type Student struct {
	ID   int
	Name string
}
