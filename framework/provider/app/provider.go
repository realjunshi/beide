package app

import (
	"beide/framework"
	"beide/framework/contract"
)

// BeideAppProvider 提供App的具体实现方法
type BeideAppProvider struct {
	BaseFolder string
}

// Register 注册HadeApp方法
func (h *BeideAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewBeideApp
}

// Boot 启动调用
func (h *BeideAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *BeideAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *BeideAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *BeideAppProvider) Name() string {
	return contract.AppKey
}
