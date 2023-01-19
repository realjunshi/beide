package app

import (
	"beide/framework"
	"beide/framework/util"
	"errors"
	"github.com/google/uuid"
	"path/filepath"
)

// BeideApp 代表hade框架的App实现
type BeideApp struct {
	// 服务容器
	container framework.Container
	// 基础路径
	baseFolder string

	appId string // 表示当前这个app的唯一id, 可以用于分布式锁等

	configMap map[string]string // 配置加载
}

// AppID 表示这个App的唯一ID
func (app BeideApp) AppID() string {
	return app.appId
}

// Version 实现版本
func (beide BeideApp) Version() string {
	return "0.0.3"
}

func (beide BeideApp) BaseFolder() string {
	if beide.baseFolder != "" {
		return beide.baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder 表示配置文件地址
func (beide BeideApp) ConfigFolder() string {
	if val, ok := beide.configMap["config_folder"]; ok {
		return val
	}
	return filepath.Join(beide.baseFolder, "config")
}

// LogFolder 表示日志存放地址
func (beide BeideApp) LogFolder() string {
	if val, ok := beide.configMap["log_folder"]; ok {
		return val
	}
	return filepath.Join(beide.StorageFolder(), "log")
}

func (beide BeideApp) HttpFolder() string {
	if val, ok := beide.configMap["http_folder"]; ok {
		return val
	}
	return filepath.Join(beide.BaseFolder(), "http")
}

func (beide BeideApp) ConsoleFolder() string {
	if val, ok := beide.configMap["console_folder"]; ok {
		return val
	}
	return filepath.Join(beide.BaseFolder(), "console")
}

func (beide BeideApp) StorageFolder() string {
	if val, ok := beide.configMap["storage_folder"]; ok {
		return val
	}
	return filepath.Join(beide.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (beide BeideApp) ProviderFolder() string {
	if val, ok := beide.configMap["provider_folder"]; ok {
		return val
	}
	return filepath.Join(beide.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (beide BeideApp) MiddlewareFolder() string {
	if val, ok := beide.configMap["middleware_folder"]; ok {
		return val
	}
	return filepath.Join(beide.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (beide BeideApp) CommandFolder() string {
	if val, ok := beide.configMap["command_folder"]; ok {
		return val
	}
	return filepath.Join(beide.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (beide BeideApp) RuntimeFolder() string {
	if val, ok := beide.configMap["runtime_folder"]; ok {
		return val
	}
	return filepath.Join(beide.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (beide BeideApp) TestFolder() string {
	if val, ok := beide.configMap["test_folder"]; ok {
		return val
	}
	return filepath.Join(beide.BaseFolder(), "test")
}

func NewBeideApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	appId := uuid.New().String()
	configMap := map[string]string{}
	return &BeideApp{baseFolder: baseFolder, container: container, appId: appId, configMap: configMap}, nil
}

// LoadAppConfig 加载配置map
func (beide *BeideApp) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		beide.configMap[key] = val
	}
}

// AppFolder 代表app目录
func (app *BeideApp) AppFolder() string {
	if val, ok := app.configMap["app_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app")
}
