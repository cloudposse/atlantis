// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
)

func AnySliceOfModelsProjectCommandContext() []models.ProjectCommandContext {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]models.ProjectCommandContext))(nil)).Elem()))
	var nullValue []models.ProjectCommandContext
	return nullValue
}

func EqSliceOfModelsProjectCommandContext(value []models.ProjectCommandContext) []models.ProjectCommandContext {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []models.ProjectCommandContext
	return nullValue
}
