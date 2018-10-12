package matchers

import (
	"reflect"

	models "github.com/cloudposse/atlantis/server/events/models"
	"github.com/petergtz/pegomock"
)

func AnyModelsPullRequestEventType() models.PullRequestEventType {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.PullRequestEventType))(nil)).Elem()))
	var nullValue models.PullRequestEventType
	return nullValue
}

func EqModelsPullRequestEventType(value models.PullRequestEventType) models.PullRequestEventType {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.PullRequestEventType
	return nullValue
}
