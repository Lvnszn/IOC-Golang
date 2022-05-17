//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by ioc-go-cli

package mysql

import (
	"github.com/alibaba/IOC-Golang/autowire"
	"github.com/alibaba/IOC-Golang/autowire/normal"
)

func init() {
	normal.RegisterStructDescriber(&autowire.StructDescriber{
		Interface: new(Mysql),
		Factory: func() interface{} {
			return &Impl{}
		},
		ParamFactory: func() interface{} {
			var _ ConfigInterface = &Config{}
			return &Config{}
		},
		ParamLoader: &paramLoader{},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(ConfigInterface)
			impl := i.(*Impl)
			return param.New(impl)
		},
	})
}

type ConfigInterface interface {
	New(impl *Impl) (*Impl, error)
}
