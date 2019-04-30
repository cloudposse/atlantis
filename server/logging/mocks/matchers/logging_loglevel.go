// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	logging "github.com/cloudposse/atlantis/server/logging"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyLoggingLogLevel() logging.LogLevel {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(logging.LogLevel))(nil)).Elem()))
	var nullValue logging.LogLevel
	return nullValue
}

func EqLoggingLogLevel(value logging.LogLevel) logging.LogLevel {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue logging.LogLevel
	return nullValue
}
