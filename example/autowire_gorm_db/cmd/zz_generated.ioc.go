//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by ioc-go-cli

package main

import (
	autowire "github.com/alibaba/IOC-Golang/autowire"
	"github.com/alibaba/IOC-Golang/autowire/singleton"
)

func init() {
	singleton.RegisterStructDescriber(&autowire.StructDescriber{
		Interface: &App{},
		Factory: func() interface{} {
			return &App{}
		},
	})
}
