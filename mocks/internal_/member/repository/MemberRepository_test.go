// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	model "github.com/cho8833/duary_lambda/internal/member/model"
	mock "github.com/stretchr/testify/mock"
)

// MemberRepository is an autogenerated mock type for the MemberRepository type
type MemberRepository struct {
	mock.Mock
}

type MemberRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MemberRepository) EXPECT() *MemberRepository_Expecter {
	return &MemberRepository_Expecter{mock: &_m.Mock}
}

// FindBySocialIdAndProvider provides a mock function with given fields: socialId, provider
func (_m *MemberRepository) FindBySocialIdAndProvider(socialId int64, provider string) (*model.Member, error) {
	ret := _m.Called(socialId, provider)

	if len(ret) == 0 {
		panic("no return value specified for FindBySocialIdAndProvider")
	}

	var r0 *model.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, string) (*model.Member, error)); ok {
		return rf(socialId, provider)
	}
	if rf, ok := ret.Get(0).(func(int64, string) *model.Member); ok {
		r0 = rf(socialId, provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(socialId, provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MemberRepository_FindBySocialIdAndProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBySocialIdAndProvider'
type MemberRepository_FindBySocialIdAndProvider_Call struct {
	*mock.Call
}

// FindBySocialIdAndProvider is a helper method to define mock.On call
//   - socialId int64
//   - provider string
func (_e *MemberRepository_Expecter) FindBySocialIdAndProvider(socialId interface{}, provider interface{}) *MemberRepository_FindBySocialIdAndProvider_Call {
	return &MemberRepository_FindBySocialIdAndProvider_Call{Call: _e.mock.On("FindBySocialIdAndProvider", socialId, provider)}
}

func (_c *MemberRepository_FindBySocialIdAndProvider_Call) Run(run func(socialId int64, provider string)) *MemberRepository_FindBySocialIdAndProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(string))
	})
	return _c
}

func (_c *MemberRepository_FindBySocialIdAndProvider_Call) Return(_a0 *model.Member, _a1 error) *MemberRepository_FindBySocialIdAndProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MemberRepository_FindBySocialIdAndProvider_Call) RunAndReturn(run func(int64, string) (*model.Member, error)) *MemberRepository_FindBySocialIdAndProvider_Call {
	_c.Call.Return(run)
	return _c
}

// SaveMember provides a mock function with given fields: member
func (_m *MemberRepository) SaveMember(member *model.Member) error {
	ret := _m.Called(member)

	if len(ret) == 0 {
		panic("no return value specified for SaveMember")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Member) error); ok {
		r0 = rf(member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MemberRepository_SaveMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveMember'
type MemberRepository_SaveMember_Call struct {
	*mock.Call
}

// SaveMember is a helper method to define mock.On call
//   - member *model.Member
func (_e *MemberRepository_Expecter) SaveMember(member interface{}) *MemberRepository_SaveMember_Call {
	return &MemberRepository_SaveMember_Call{Call: _e.mock.On("SaveMember", member)}
}

func (_c *MemberRepository_SaveMember_Call) Run(run func(member *model.Member)) *MemberRepository_SaveMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Member))
	})
	return _c
}

func (_c *MemberRepository_SaveMember_Call) Return(_a0 error) *MemberRepository_SaveMember_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MemberRepository_SaveMember_Call) RunAndReturn(run func(*model.Member) error) *MemberRepository_SaveMember_Call {
	_c.Call.Return(run)
	return _c
}

// NewMemberRepository creates a new instance of MemberRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMemberRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MemberRepository {
	mock := &MemberRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
