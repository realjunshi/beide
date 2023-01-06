package framework

// IGroup 代表前缀分组
type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

// Group struct 实现了IGroup

type Group struct {
	core   *Core
	prefix string
}

//初始化 Group

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}

// Get 实现Get方法
func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Get(uri, handler)
}

// Post 实现Post方法
func (g *Group) Post(uri string, handler ControllerHandler) {
	//uri = g.getAbsolutePrefix() + uri
	uri = g.prefix + uri
	g.core.Post(uri, handler)
}

// Put 实现Put方法
func (g *Group) Put(uri string, handler ControllerHandler) {
	//uri = g.getAbsolutePrefix() + uri
	uri = g.prefix + uri
	g.core.Put(uri, handler)
}

// Delete 实现Delete方法
func (g *Group) Delete(uri string, handler ControllerHandler) {
	//uri = g.getAbsolutePrefix() + uri
	uri = g.prefix + uri
	g.core.Delete(uri, handler)
}

// 获取当前group的绝对路径
//func (g *Group) getAbsolutePrefix() string {
//	if g.parent == nil {
//		return g.prefix
//	}
//	return g.parent.getAbsolutePrefix() + g.prefix
//}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
