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

	ioc "github.com/alibaba/IOC-Golang"
	"github.com/alibaba/IOC-Golang/autowire/singleton"

	normalMysql "github.com/alibaba/IOC-Golang/extension/normal/mysql"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type App struct {
	MyDataTable normalMysql.Mysql `normal:"Impl,my-mysql,mydata"`
}

type MyDataDO struct {
	Id    int32
	Value string
}

func (a *MyDataDO) TableName() string {
	return "mydata"
}

func (a *App) Run() {
	myDataDOs := make([]MyDataDO, 0)
	if err := a.MyDataTable.SelectWhere("id = ?", &myDataDOs, 1); err != nil {
		panic(err)
	}
	fmt.Println(myDataDOs)
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
