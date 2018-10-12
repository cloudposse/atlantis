package matchers

import (
	"reflect"

	events "github.com/cloudposse/atlantis/server/events"
	"github.com/petergtz/pegomock"
)

func AnyPtrToEventsCommentCommand() *events.CommentCommand {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*events.CommentCommand))(nil)).Elem()))
	var nullValue *events.CommentCommand
	return nullValue
}

func EqPtrToEventsCommentCommand(value *events.CommentCommand) *events.CommentCommand {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *events.CommentCommand
	return nullValue
}
