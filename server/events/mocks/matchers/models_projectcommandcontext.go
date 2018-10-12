package matchers

import (
	"reflect"

	models "github.com/cloudposse/atlantis/server/events/models"
	"github.com/petergtz/pegomock"
)

func AnyModelsProjectCommandContext() models.ProjectCommandContext {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.ProjectCommandContext))(nil)).Elem()))
	var nullValue models.ProjectCommandContext
	return nullValue
}

func EqModelsProjectCommandContext(value models.ProjectCommandContext) models.ProjectCommandContext {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.ProjectCommandContext
	return nullValue
}
