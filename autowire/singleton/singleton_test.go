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

package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alibaba/ioc-golang/autowire"
)

type mockInterface interface {
}

type mockImpl struct {
}

const mockInterfaceName = "mockInterface"
const mockImplName = "mockImpl"

func TestAutowire_RegisterAndGetAllStructDescribers(t *testing.T) {
	t.Run("test singleton autowire register and get all struct describers", func(t *testing.T) {
		sd := &autowire.StructDescriber{
			Interface: new(mockInterface),
			Factory: func() interface{} {
				return &mockImpl{}
			},
		}
		RegisterStructDescriber(sd)
		n := &SingletonAutowire{}
		allStructDesc := n.GetAllStructDescribers()
		assert.NotNil(t, allStructDesc)
		structDesc, ok := allStructDesc[mockInterfaceName+"-"+mockImplName]
		assert.True(t, ok)
		assert.Equal(t, mockInterfaceName+"-"+mockImplName, structDesc.ID())
	})
}

func TestAutowire_TagKey(t *testing.T) {
	t.Run("test singleton autowire tag", func(t *testing.T) {
		n := &SingletonAutowire{}
		assert.Equal(t, Name, n.TagKey())
	})
}

func TestAutowire_IsSingleton(t *testing.T) {
	t.Run("test singleton autowire isSingleton", func(t *testing.T) {
		n := &SingletonAutowire{}
		assert.Equal(t, true, n.IsSingleton())
	})
}

func TestAutowire_CanBeEntrance(t *testing.T) {
	t.Run("test normal autowire can be entrance", func(t *testing.T) {
		n := &SingletonAutowire{}
		assert.Equal(t, true, n.CanBeEntrance())
	})
}

func TestNewSingletonAutowire(t *testing.T) {
	singletonAutowire := NewSingletonAutowire(nil, nil, nil)
	assert.NotNil(t, singletonAutowire)
}
