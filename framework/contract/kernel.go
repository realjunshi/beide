package contract

import "net/http"

const KernelKey = "beide:kernel"

// Kernel 接口提供框架最核心的结构
type Kernel interface {
	HttpEngine() http.Handler
}
