package matchers

import (
	"reflect"

	valid "github.com/cloudposse/atlantis/server/events/yaml/valid"
	"github.com/petergtz/pegomock"
)

func AnyValidConfig() valid.Config {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(valid.Config))(nil)).Elem()))
	var nullValue valid.Config
	return nullValue
}

func EqValidConfig(value valid.Config) valid.Config {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue valid.Config
	return nullValue
}
