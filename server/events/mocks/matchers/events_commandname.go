package matchers

import (
	"reflect"

	events "github.com/cloudposse/atlantis/server/events"
	"github.com/petergtz/pegomock"
)

func AnyEventsCommandName() events.CommandName {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.CommandName))(nil)).Elem()))
	var nullValue events.CommandName
	return nullValue
}

func EqEventsCommandName(value events.CommandName) events.CommandName {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.CommandName
	return nullValue
}
