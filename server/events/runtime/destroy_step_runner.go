package runtime

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudposse/atlantis/server/events/models"
	"github.com/hashicorp/go-version"
)

// DestroyStepRunner runs `terraform destroy`.
type DestroyStepRunner struct {
	TerraformExecutor TerraformExec
}

func (a *DestroyStepRunner) Run(ctx models.ProjectCommandContext, extraArgs []string, path string) (string, error) {
	planPath := filepath.Join(path, GetPlanFilename(ctx.Workspace, ctx.ProjectConfig))
	stat, err := os.Stat(planPath)
	if err != nil || stat.IsDir() {
		return "", fmt.Errorf("no plan found at path %q and workspace %q – did you run plan and apply?", ctx.RepoRelDir, ctx.Workspace)
	}

	// NOTE: we need to quote the plan path because Bitbucket Server can
	// have spaces in its repo owner names which is part of the path.
	tfDestroyCmd := append(append(append([]string{"destroy", "-input=false", "-no-color", "-auto-approve"}, extraArgs...), ctx.CommentArgs...), fmt.Sprintf("%q", planPath))
	var tfVersion *version.Version
	if ctx.ProjectConfig != nil && ctx.ProjectConfig.TerraformVersion != nil {
		tfVersion = ctx.ProjectConfig.TerraformVersion
	}
	out, tfErr := a.TerraformExecutor.RunCommandWithVersion(ctx.Log, path, tfDestroyCmd, tfVersion, ctx.Workspace)

	if tfErr == nil {
		ctx.Log.Info("destroy successful")
	}
	return out, tfErr
}
