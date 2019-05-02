package raw

import (
	"github.com/cloudposse/atlantis/server/events/yaml/valid"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Stage struct {
	Steps []Step `yaml:"steps,omitempty" json:"steps,omitempty"`
}

func (s Stage) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Steps),
	)
}

func (s Stage) ToValid() valid.Stage {
	var validSteps []valid.Step
	for _, s := range s.Steps {
		validSteps = append(validSteps, s.ToValid())
	}
	return valid.Stage{
		Steps: validSteps,
	}
}
