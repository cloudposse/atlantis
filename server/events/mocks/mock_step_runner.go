// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: StepRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockStepRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockStepRunner(options ...pegomock.Option) *MockStepRunner {
	mock := &MockStepRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockStepRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockStepRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockStepRunner) Run(ctx models.ProjectCommandContext, extraArgs []string, path string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockStepRunner().")
	}
	params := []pegomock.Param{ctx, extraArgs, path}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockStepRunner) VerifyWasCalledOnce() *VerifierStepRunner {
	return &VerifierStepRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockStepRunner) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierStepRunner {
	return &VerifierStepRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockStepRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierStepRunner {
	return &VerifierStepRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockStepRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierStepRunner {
	return &VerifierStepRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierStepRunner struct {
	mock                   *MockStepRunner
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierStepRunner) Run(ctx models.ProjectCommandContext, extraArgs []string, path string) *StepRunner_Run_OngoingVerification {
	params := []pegomock.Param{ctx, extraArgs, path}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params, verifier.timeout)
	return &StepRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type StepRunner_Run_OngoingVerification struct {
	mock              *MockStepRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *StepRunner_Run_OngoingVerification) GetCapturedArguments() (models.ProjectCommandContext, []string, string) {
	ctx, extraArgs, path := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], extraArgs[len(extraArgs)-1], path[len(path)-1]
}

func (c *StepRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext, _param1 [][]string, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
		_param1 = make([][]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]string)
		}
		_param2 = make([]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}
