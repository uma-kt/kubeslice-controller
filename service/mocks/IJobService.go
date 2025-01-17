// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// IJobService is an autogenerated mock type for the IJobService type
type IJobService struct {
	mock.Mock
}

// CreateJob provides a mock function with given fields: ctx, namespace, jobImage, environment
func (_m *IJobService) CreateJob(ctx context.Context, namespace string, jobImage string, environment map[string]string) (reconcile.Result, error) {
	ret := _m.Called(ctx, namespace, jobImage, environment)

	var r0 reconcile.Result
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]string) reconcile.Result); ok {
		r0 = rf(ctx, namespace, jobImage, environment)
	} else {
		r0 = ret.Get(0).(reconcile.Result)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, map[string]string) error); ok {
		r1 = rf(ctx, namespace, jobImage, environment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
