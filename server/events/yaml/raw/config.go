package raw

import (
	"errors"

	"github.com/cloudposse/atlantis/server/events/yaml/valid"
	"github.com/go-ozzo/ozzo-validation"
)

// Config is the representation for the whole config file at the top level.
type Config struct {
	Version   *int                `yaml:"version,omitempty"`
	Projects  []Project           `yaml:"projects,omitempty"`
	Workflows map[string]Workflow `yaml:"workflows,omitempty"`
}

func (c Config) Validate() error {
	equals2 := func(value interface{}) error {
		asIntPtr := value.(*int)
		if asIntPtr == nil {
			return errors.New("is required. If you've just upgraded Atlantis you need to rewrite your atlantis.yaml for version 2. See www.runatlantis.io/docs/upgrading-atlantis-yaml-to-version-2.html")
		}
		if *asIntPtr != 2 {
			return errors.New("must equal 2")
		}
		return nil
	}
	return validation.ValidateStruct(&c,
		validation.Field(&c.Version, validation.By(equals2)),
		validation.Field(&c.Projects),
		validation.Field(&c.Workflows),
	)
}

func (c Config) ToValid() valid.Config {
	var validProjects []valid.Project
	for _, p := range c.Projects {
		validProjects = append(validProjects, p.ToValid())
	}

	validWorkflows := make(map[string]valid.Workflow)
	for k, v := range c.Workflows {
		validWorkflows[k] = v.ToValid()
	}
	return valid.Config{
		Version:   *c.Version,
		Projects:  validProjects,
		Workflows: validWorkflows,
	}
}
