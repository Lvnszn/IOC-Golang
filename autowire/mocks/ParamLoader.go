// Code generated by mockery v2.12.2. DO NOT EDIT.

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

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	autowire "github.com/alibaba/IOC-Golang/autowire"
)

// ParamLoader is an autogenerated mock type for the ParamLoader type
type ParamLoader struct {
	mock.Mock
}

// Load provides a mock function with given fields: sd, fi
func (_m *ParamLoader) Load(sd *autowire.StructDescriber, fi *autowire.FieldInfo) (interface{}, error) {
	ret := _m.Called(sd, fi)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*autowire.StructDescriber, *autowire.FieldInfo) interface{}); ok {
		r0 = rf(sd, fi)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*autowire.StructDescriber, *autowire.FieldInfo) error); ok {
		r1 = rf(sd, fi)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewParamLoader creates a new instance of ParamLoader. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewParamLoader(t testing.TB) *ParamLoader {
	mock := &ParamLoader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
