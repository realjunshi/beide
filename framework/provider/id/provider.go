package id

import (
  "beide/framework"
  "beide/framework/contract"
)

type HadeIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *HadeIDProvider) Register(c framework.Container) framework.NewInstance {
  //log.Println("HadeIDProvider:", 1234)
  return NewHadeIDService
}

// Boot will called when the service instantiate
func (provider *HadeIDProvider) Boot(c framework.Container) error {
  return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HadeIDProvider) IsDefer() bool {
  return false
}

// Params define the necessary params for NewInstance
func (provider *HadeIDProvider) Params(c framework.Container) []interface{} {
  return []interface{}{}
}

/// Name define the name for this service
func (provider *HadeIDProvider) Name() string {
  return contract.IDKey
}
