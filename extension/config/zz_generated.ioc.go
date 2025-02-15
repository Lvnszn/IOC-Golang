//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by ioc-go-cli

package config

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	autowireconfig "github.com/alibaba/ioc-golang/extension/autowire/config"
)

func init() {
	autowireconfig.RegisterStructDescriber(&autowire.StructDescriber{
		Interface: new(ConfigInt),
		Factory: func() interface{} {
			return new(ConfigInt)
		},
		ParamFactory: func() interface{} {
			return new(ConfigInt)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(ConfigIntInterface)
			impl := i.(*ConfigInt)
			return param.New(impl)
		},
	})
	autowireconfig.RegisterStructDescriber(&autowire.StructDescriber{
		Interface: new(ConfigMap),
		Factory: func() interface{} {
			return new(ConfigMap)
		},
		ParamFactory: func() interface{} {
			return new(ConfigMap)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(ConfigMapInterface)
			impl := i.(*ConfigMap)
			return param.New(impl)
		},
	})
	autowireconfig.RegisterStructDescriber(&autowire.StructDescriber{
		Interface: new(ConfigSlice),
		Factory: func() interface{} {
			return new(ConfigSlice)
		},
		ParamFactory: func() interface{} {
			return new(ConfigSlice)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(ConfigSliceInterface)
			impl := i.(*ConfigSlice)
			return param.New(impl)
		},
	})
	autowireconfig.RegisterStructDescriber(&autowire.StructDescriber{
		Interface: new(ConfigString),
		Factory: func() interface{} {
			return new(ConfigString)
		},
		ParamFactory: func() interface{} {
			return new(ConfigString)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(ConfigStringInterface)
			impl := i.(*ConfigString)
			return param.New(impl)
		},
	})
}

type ConfigIntInterface interface {
	New(impl *ConfigInt) (*ConfigInt, error)
}
type ConfigMapInterface interface {
	New(impl *ConfigMap) (*ConfigMap, error)
}
type ConfigSliceInterface interface {
	New(impl *ConfigSlice) (*ConfigSlice, error)
}
type ConfigStringInterface interface {
	New(impl *ConfigString) (*ConfigString, error)
}
