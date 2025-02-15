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

package mysql

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	ioc "github.com/alibaba/ioc-golang"
	"github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/config"
)

func Test_getMysqlLinkStr(t *testing.T) {
	conf := &Config{
		Host:     "192.168.1.1",
		Port:     "1234",
		Username: "admin",
		Password: "admin",
		DBName:   "mydb",
	}
	got := getMysqlLinkStr(conf)
	assert.Equal(t, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName), got)
}

func Test_paramLoader_Load(t *testing.T) {
	p := &paramLoader{}
	assert.Nil(t, os.Setenv(config.EnvKeyIOCGolangConfigPath, "./test/ioc_golang.yaml"))
	assert.Nil(t, ioc.Load())
	got, err := p.Load(nil, &autowire.FieldInfo{
		TagValue: "xxx,my-mysql,tableName",
	})
	assert.Nil(t, err)
	config, ok := got.(*Config)
	assert.True(t, ok)
	assert.Equal(t, "192.168.1.1", config.Host)
	assert.Equal(t, "1234", config.Port)
	assert.Equal(t, "admin", config.Username)
	assert.Equal(t, "admin", config.Password)
	assert.Equal(t, "tableName", config.TableName)
	assert.Equal(t, "mydb", config.DBName)
}
