// Code generated by mockery v2.19.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "sigs.k8s.io/controller-runtime/pkg/client"

	mock "github.com/stretchr/testify/mock"

	types "github.com/aws-controllers-k8s/runtime/pkg/types"
)

// ResolvedReferenceManager is an autogenerated mock type for the ResolvedReferenceManager type
type ResolvedReferenceManager struct {
	mock.Mock
}

// ClearResolvedReferences provides a mock function with given fields: _a0
func (_m *ResolvedReferenceManager) ClearResolvedReferences(_a0 types.AWSResource) types.AWSResource {
	ret := _m.Called(_a0)

	var r0 types.AWSResource
	if rf, ok := ret.Get(0).(func(types.AWSResource) types.AWSResource); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResource)
		}
	}

	return r0
}

// ResolveReferences provides a mock function with given fields: _a0, _a1, _a2
func (_m *ResolvedReferenceManager) ResolveReferences(_a0 context.Context, _a1 client.Reader, _a2 types.AWSResource) (types.AWSResource, bool, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 types.AWSResource
	if rf, ok := ret.Get(0).(func(context.Context, client.Reader, types.AWSResource) types.AWSResource); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResource)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, client.Reader, types.AWSResource) bool); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, client.Reader, types.AWSResource) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewResolvedReferenceManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewResolvedReferenceManager creates a new instance of ResolvedReferenceManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResolvedReferenceManager(t mockConstructorTestingTNewResolvedReferenceManager) *ResolvedReferenceManager {
	mock := &ResolvedReferenceManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
