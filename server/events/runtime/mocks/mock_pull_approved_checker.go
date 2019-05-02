// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/runtime (interfaces: PullApprovedChecker)

package mocks

import (
	models "github.com/cloudposse/atlantis/server/events/models"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockPullApprovedChecker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPullApprovedChecker(options ...pegomock.Option) *MockPullApprovedChecker {
	mock := &MockPullApprovedChecker{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPullApprovedChecker) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPullApprovedChecker) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPullApprovedChecker) PullIsApproved(baseRepo models.Repo, pull models.PullRequest) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPullApprovedChecker().")
	}
	params := []pegomock.Param{baseRepo, pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsApproved", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockPullApprovedChecker) VerifyWasCalledOnce() *VerifierPullApprovedChecker {
	return &VerifierPullApprovedChecker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPullApprovedChecker) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierPullApprovedChecker {
	return &VerifierPullApprovedChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPullApprovedChecker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierPullApprovedChecker {
	return &VerifierPullApprovedChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPullApprovedChecker) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierPullApprovedChecker {
	return &VerifierPullApprovedChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierPullApprovedChecker struct {
	mock                   *MockPullApprovedChecker
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierPullApprovedChecker) PullIsApproved(baseRepo models.Repo, pull models.PullRequest) *PullApprovedChecker_PullIsApproved_OngoingVerification {
	params := []pegomock.Param{baseRepo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsApproved", params, verifier.timeout)
	return &PullApprovedChecker_PullIsApproved_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type PullApprovedChecker_PullIsApproved_OngoingVerification struct {
	mock              *MockPullApprovedChecker
	methodInvocations []pegomock.MethodInvocation
}

func (c *PullApprovedChecker_PullIsApproved_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	baseRepo, pull := c.GetAllCapturedArguments()
	return baseRepo[len(baseRepo)-1], pull[len(pull)-1]
}

func (c *PullApprovedChecker_PullIsApproved_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}
