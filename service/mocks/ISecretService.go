// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

// ISecretService is an autogenerated mock type for the ISecretService type
type ISecretService struct {
	mock.Mock
}

// DeleteSecret provides a mock function with given fields: ctx, namespace, secretName
func (_m *ISecretService) DeleteSecret(ctx context.Context, namespace string, secretName string) (reconcile.Result, error) {
	ret := _m.Called(ctx, namespace, secretName)

	var r0 reconcile.Result
	if rf, ok := ret.Get(0).(func(context.Context, string, string) reconcile.Result); ok {
		r0 = rf(ctx, namespace, secretName)
	} else {
		r0 = ret.Get(0).(reconcile.Result)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, secretName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}