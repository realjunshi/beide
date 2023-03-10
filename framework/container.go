package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	//Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)

	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会 panic。
	//所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的 params 参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// BeideContainer 是服务容器的具体实现
type BeideContainer struct {
	Container // 强制要求 HadeContainer 实现 Container 接口
	// providers 存储注册的服务提供者，key 为字符串凭证
	providers map[string]ServiceProvider
	// instance 存储具体的实例，key 为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作
	lock *sync.RWMutex
}

// NewBeideContainer 创建一个服务容器
func NewBeideContainer() *BeideContainer {
	return &BeideContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      &sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (beide *BeideContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range beide.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind 将服务容器和关键字做了绑定
func (beide *BeideContainer) Bind(provider ServiceProvider) error {
	//beide.lock.Lock()
	//defer beide.lock.Unlock()
	key := provider.Name()

	//log.Println("provider name:", key)
	//log.Println("provider.IsDefer() :", provider.IsDefer())
	beide.providers[key] = provider
	//beide.PrintProviders()

	// if provider is not defer
	if provider.IsDefer() == false {
		if err := provider.Boot(beide); err != nil {
			return err
		}

		// 实例化方法
		params := provider.Params(beide)
		method := provider.Register(beide)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		beide.instances[key] = instance
	}
	return nil
}

// Make 方式调用内部的 make 实现
func (beide *BeideContainer) Make(key string) (interface{}, error) {
	return beide.make(key, nil, false)
}

func (beide *BeideContainer) MustMake(key string) interface{} {
	serv, err := beide.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

// MakeNew 方式使用内部的 make 初始化
func (beide *BeideContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return beide.make(key, nil, false)
}

// IsBind 是否已经绑定了
func (beide *BeideContainer) IsBind(key string) bool {
	return beide.findServiceProvider(key) != nil
}

// newInstance
func (beide *BeideContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(beide); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(beide)
	}
	method := sp.Register(beide)
	ins, err := method(params)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// 根据key查找服务发现者
func (beide *BeideContainer) findServiceProvider(key string) ServiceProvider {
	beide.lock.RLock()
	defer beide.lock.RUnlock()
	if sp, ok := beide.providers[key]; ok {
		return sp
	}
	return nil
}

// 真正的实例化一个服务
func (beide *BeideContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	beide.lock.RLock()
	defer beide.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := beide.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return beide.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := beide.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := beide.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	beide.instances[key] = inst
	return inst, nil
}

// NameList 列出容器中所有服务提供者的字符串凭证
func (beide *BeideContainer) NameList() []string {
	ret := []string{}
	for _, provider := range beide.providers {
		name := provider.Name()
		ret = append(ret, name)
	}
	return ret
}
