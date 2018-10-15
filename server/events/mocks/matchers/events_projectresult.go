package matchers

import (
	"reflect"

	events "github.com/cloudposse/atlantis/server/events"
	"github.com/petergtz/pegomock"
)

func AnyEventsProjectResult() events.ProjectResult {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.ProjectResult))(nil)).Elem()))
	var nullValue events.ProjectResult
	return nullValue
}

func EqEventsProjectResult(value events.ProjectResult) events.ProjectResult {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.ProjectResult
	return nullValue
}
