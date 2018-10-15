package matchers

import (
	"reflect"

	models "github.com/cloudposse/atlantis/server/events/models"
	"github.com/petergtz/pegomock"
)

func AnyPtrToModelsPullRequest() *models.PullRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*models.PullRequest))(nil)).Elem()))
	var nullValue *models.PullRequest
	return nullValue
}

func EqPtrToModelsPullRequest(value *models.PullRequest) *models.PullRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *models.PullRequest
	return nullValue
}
