package framework

import (
	"log"
	"net/http"
	"strings"
)

// 建立请求与对应控制器的映射关系

type Core struct {
	//router map[string]map[string]ControllerHandler
	router map[string]*Tree
}

func NewCore() *Core {

	// 定义二级map
	//getRouter := map[string]*Tree{}
	//postRouter := map[string]*Tree{}
	//putRouter := map[string]*Tree{}
	//deleteRouter := map[string]*Tree{}

	// 将二级map写入一级map
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{
		router: router,
	}
}

// Get 对应 Method = Get
func (c *Core) Get(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["GET"][upperUrl] = handler
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Post 对应 Method = POST
func (c *Core) Post(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["POST"][upperUrl] = handler
	if err := c.router["Post"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Put 对应 Method = PUT
func (c *Core) Put(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["Put"][upperUrl] = handler
	if err := c.router["Put"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Delete 对应 Method = DELETE
func (c *Core) Delete(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["DELETE"][upperUrl] = handler
	if err := c.router["Delete"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// FindRouteByRequest 匹配路由
func (c *Core) FindRouteByRequest(r *http.Request) ControllerHandler {
	// 请求URL中解析参数
	uri := r.URL.Path
	method := r.Method
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	router := c.FindRouteByRequest(request)
	if router == nil {
		// 如果没有找到，这里打印日志
		err := ctx.Json(404, "not found")
		if err != nil {
			return
		}
	}

	// 调用路由函数，如果返回err代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		err := ctx.Json(500, "inner error")
		if err != nil {
			return
		}
	}
	log.Println("core.router")
}
