package events

import (
	"github.com/cloudposse/atlantis/server/events/models"
	"github.com/pkg/errors"
)

// BuildDestroyCommands builds project custom commands for this comment. If the
// comment doesn't specify one project then there may be multiple commands
// to be run.
func (p *DefaultProjectCommandBuilder) BuildCustomCommands(ctx *CommandContext, cmd *CommentCommand) ([]models.ProjectCommandContext, error) {
	if !cmd.IsForSpecificProject() {
		return p.buildCustomAllCommands(ctx, cmd)
	}
	pac, err := p.buildProjectCustomCommand(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return []models.ProjectCommandContext{pac}, nil
}

func (p *DefaultProjectCommandBuilder) buildCustomAllCommands(ctx *CommandContext, commentCmd *CommentCommand) ([]models.ProjectCommandContext, error) {
	// lock all dirs in this pull request
	unlockFn, err := p.WorkingDirLocker.TryLockPull(ctx.BaseRepo.FullName, ctx.Pull.Num)
	if err != nil {
		return nil, err
	}
	defer unlockFn()

	pullDir, err := p.WorkingDir.GetPullDir(ctx.BaseRepo, ctx.Pull)
	if err != nil {
		return nil, err
	}

	plans, err := p.PendingPlanFinder.Find(pullDir)
	if err != nil {
		return nil, err
	}

	var cmds []models.ProjectCommandContext
	for _, plan := range plans {
		cmd, err := p.buildProjectCommandCtx(ctx, commentCmd.ProjectName, commentCmd.Flags, plan.RepoDir, plan.RepoRelDir, plan.Workspace)
		if err != nil {
			return nil, errors.Wrapf(err, "building custom command for dir %q", plan.RepoRelDir)
		}
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}

func (p *DefaultProjectCommandBuilder) buildProjectCustomCommand(ctx *CommandContext, cmd *CommentCommand) (models.ProjectCommandContext, error) {
	workspace := DefaultWorkspace
	if cmd.Workspace != "" {
		workspace = cmd.Workspace
	}

	var projCtx models.ProjectCommandContext
	unlockFn, err := p.WorkingDirLocker.TryLock(ctx.BaseRepo.FullName, ctx.Pull.Num, workspace)
	if err != nil {
		return projCtx, err
	}
	defer unlockFn()

	repoDir, err := p.WorkingDir.GetWorkingDir(ctx.BaseRepo, ctx.Pull, workspace)
	if err != nil {
		return projCtx, err
	}

	repoRelDir := DefaultRepoRelDir
	if cmd.RepoRelDir != "" {
		repoRelDir = cmd.RepoRelDir
	}

	return p.buildProjectCommandCtx(ctx, cmd.ProjectName, cmd.Flags, repoDir, repoRelDir, workspace)
}
