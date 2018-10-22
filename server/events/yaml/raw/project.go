package raw

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cloudposse/atlantis/server/events/yaml/valid"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/hashicorp/go-version"
	"github.com/pkg/errors"
)

const (
	DefaultWorkspace           = "default"
	ApprovedApplyRequirement   = "approved"
	ApprovedDestroyRequirement = "approved"
)

type Project struct {
	Name                *string   `yaml:"name,omitempty"`
	Dir                 *string   `yaml:"dir,omitempty"`
	Workspace           *string   `yaml:"workspace,omitempty"`
	Workflow            *string   `yaml:"workflow,omitempty"`
	TerraformVersion    *string   `yaml:"terraform_version,omitempty"`
	Autoplan            *Autoplan `yaml:"autoplan,omitempty"`
	ApplyRequirements   []string  `yaml:"apply_requirements,omitempty"`
	DestroyRequirements []string  `yaml:"destroy_requirements,omitempty"`
}

func (p Project) Validate() error {
	hasDotDot := func(value interface{}) error {
		if strings.Contains(*value.(*string), "..") {
			return errors.New("cannot contain '..'")
		}
		return nil
	}
	validApplyReq := func(value interface{}) error {
		reqs := value.([]string)
		for _, r := range reqs {
			if r != ApprovedApplyRequirement {
				return fmt.Errorf("%q not supported, only %s is supported", r, ApprovedApplyRequirement)
			}
		}
		return nil
	}
	validDestroyReq := func(value interface{}) error {
		reqs := value.([]string)
		for _, r := range reqs {
			if r != ApprovedDestroyRequirement {
				return fmt.Errorf("%q not supported, only %s is supported", r, ApprovedDestroyRequirement)
			}
		}
		return nil
	}
	validTFVersion := func(value interface{}) error {
		strPtr := value.(*string)
		if strPtr == nil {
			return nil
		}
		_, err := version.NewVersion(*strPtr)
		return errors.Wrapf(err, "version %q could not be parsed", *strPtr)
	}
	validName := func(value interface{}) error {
		strPtr := value.(*string)
		if strPtr == nil {
			return nil
		}
		if *strPtr == "" {
			return errors.New("if set cannot be empty")
		}
		return nil
	}
	return validation.ValidateStruct(&p,
		validation.Field(&p.Dir, validation.Required, validation.By(hasDotDot)),
		validation.Field(&p.ApplyRequirements, validation.By(validApplyReq)),
		validation.Field(&p.DestroyRequirements, validation.By(validDestroyReq)),
		validation.Field(&p.TerraformVersion, validation.By(validTFVersion)),
		validation.Field(&p.Name, validation.By(validName)),
	)
}

func (p Project) ToValid() valid.Project {
	var v valid.Project
	cleanedDir := filepath.Clean(*p.Dir)
	if cleanedDir == "/" {
		cleanedDir = "."
	}
	v.Dir = cleanedDir

	if p.Workspace == nil || *p.Workspace == "" {
		v.Workspace = DefaultWorkspace
	} else {
		v.Workspace = *p.Workspace
	}

	v.Workflow = p.Workflow
	if p.TerraformVersion != nil {
		v.TerraformVersion, _ = version.NewVersion(*p.TerraformVersion)
	}
	if p.Autoplan == nil {
		v.Autoplan = DefaultAutoPlan()
	} else {
		v.Autoplan = p.Autoplan.ToValid()
	}

	// There are no default apply requirements.
	v.ApplyRequirements = p.ApplyRequirements

	// There are no default destroy requirements.
	v.DestroyRequirements = p.DestroyRequirements

	v.Name = p.Name

	return v
}
