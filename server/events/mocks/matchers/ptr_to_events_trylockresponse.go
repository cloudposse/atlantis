package matchers

import (
	"reflect"

	events "github.com/cloudposse/atlantis/server/events"
	"github.com/petergtz/pegomock"
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
