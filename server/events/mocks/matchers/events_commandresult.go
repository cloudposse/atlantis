package matchers

import (
	"reflect"

	events "github.com/cloudposse/atlantis/server/events"
	"github.com/petergtz/pegomock"
)

func AnyEventsCommandResult() events.CommandResult {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.CommandResult))(nil)).Elem()))
	var nullValue events.CommandResult
	return nullValue
}

func EqEventsCommandResult(value events.CommandResult) events.CommandResult {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.CommandResult
	return nullValue
}
