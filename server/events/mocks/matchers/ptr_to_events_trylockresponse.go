// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	events "github.com/cloudposse/atlantis/server/events"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyPtrToEventsTryLockResponse() *events.TryLockResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*events.TryLockResponse))(nil)).Elem()))
	var nullValue *events.TryLockResponse
	return nullValue
}

func EqPtrToEventsTryLockResponse(value *events.TryLockResponse) *events.TryLockResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *events.TryLockResponse
	return nullValue
}
