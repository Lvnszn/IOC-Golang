/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	"github.com/go-redis/redis"

	ioc "github.com/alibaba/ioc-golang"
	"github.com/alibaba/ioc-golang/autowire/singleton"

	normalRedis "github.com/alibaba/ioc-golang/extension/normal/redis"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:paramType=Param
// +ioc:autowire:constructFunc=Init

type App struct {
	NormalRedis    normalRedis.Redis `normal:"Impl"`
	NormalDB1Redis normalRedis.Redis `normal:"Impl,db1-redis"`
	NormalDB2Redis normalRedis.Redis `normal:"Impl,db2-redis"`
	NormalDB3Redis normalRedis.Redis `normal:"Impl,address=127.0.0.1:6379&db=3"`

	privateClient *redis.Client
}

type Param struct {
	RedisAddr string
}

func (p *Param) Init(a *App) (*App, error) {
	privateClient := redis.NewClient(&redis.Options{
		Addr: p.RedisAddr,
	})
	a.privateClient = privateClient
	return a, nil
}

func (a *App) Run() {
	if _, err := a.NormalRedis.Set("mykey", "db0", -1); err != nil {
		panic(err)
	}

	if _, err := a.NormalDB1Redis.Set("mykey", "db1", -1); err != nil {
		panic(err)
	}

	if _, err := a.NormalDB2Redis.Set("mykey", "db2", -1); err != nil {
		panic(err)
	}

	if _, err := a.NormalDB3Redis.Set("mykey", "db3", -1); err != nil {
		panic(err)
	}

	val1, err := a.NormalRedis.Get("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println("client0 get ", val1)

	val2, err := a.NormalDB1Redis.Get("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println("client1 get ", val2)

	val3, err := a.NormalDB2Redis.Get("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println("client2 get ", val3)

	val4, err := a.NormalDB3Redis.Get("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println("client3 get ", val4)
}

func main() {
	if err := ioc.Load(); err != nil {
		panic(err)
	}
	appInterface, err := singleton.GetImpl("App-App")
	if err != nil {
		panic(err)
	}
	app := appInterface.(*App)

	app.Run()
}
