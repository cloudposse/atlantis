// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	models "github.com/cloudposse/atlantis/server/events/models"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyModelsPullRequest() models.PullRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.PullRequest))(nil)).Elem()))
	var nullValue models.PullRequest
	return nullValue
}

func EqModelsPullRequest(value models.PullRequest) models.PullRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.PullRequest
	return nullValue
}
