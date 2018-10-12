package matchers

import (
	"reflect"

	models "github.com/cloudposse/atlantis/server/events/models"
	"github.com/petergtz/pegomock"
)

func AnyModelsVCSHostType() models.VCSHostType {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.VCSHostType))(nil)).Elem()))
	var nullValue models.VCSHostType
	return nullValue
}

func EqModelsVCSHostType(value models.VCSHostType) models.VCSHostType {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.VCSHostType
	return nullValue
}
