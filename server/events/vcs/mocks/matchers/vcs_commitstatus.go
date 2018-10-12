package matchers

import (
	"reflect"

	models "github.com/cloudposse/atlantis/server/events/models"
	"github.com/petergtz/pegomock"
)

func AnyVcsCommitStatus() models.CommitStatus {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.CommitStatus))(nil)).Elem()))
	var nullValue models.CommitStatus
	return nullValue
}

func EqVcsCommitStatus(value models.CommitStatus) models.CommitStatus {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.CommitStatus
	return nullValue
}
