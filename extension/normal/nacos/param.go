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

package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type Config struct {
	vo.NacosClientParam
}

func (c *Config) New(impl *Impl) (*Impl, error) {
	var err error
	impl.IConfigClient, err = clients.NewConfigClient(c.NacosClientParam)
	if err != nil {
		return nil, err
	}
	impl.INamingClient, err = clients.NewNamingClient(c.NacosClientParam)
	if err != nil {
		return nil, err
	}
	return impl, nil
}
