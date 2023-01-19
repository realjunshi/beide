package env

import (
  "beide/framework"
  "beide/framework/contract"
)

type HadeEnvProvider struct {
  Folder string
}

// Register registe a new function for make a service instance
func (provider *HadeEnvProvider) Register(c framework.Container) framework.NewInstance {
  //log.Println("Register HadeEnvProvider:", 1234567)
  return NewHadeEnv
}

// Boot will called when the service instantiate
func (provider *HadeEnvProvider) Boot(c framework.Container) error {
  //log.Println("HadeEnvProvider boot", contract.AppKey)
  app := c.MustMake(contract.AppKey).(contract.App)
  //log.Println("app:Boot ", app)
  provider.Folder = app.BaseFolder()
  return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HadeEnvProvider) IsDefer() bool {
  return false
}

// Params define the necessary params for NewInstance
func (provider *HadeEnvProvider) Params(c framework.Container) []interface{} {
  //log.Println("HadeEnvProvider:", provider.Folder)
  return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *HadeEnvProvider) Name() string {
  return contract.EnvKey
}
