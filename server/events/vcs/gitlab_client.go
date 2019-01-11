// Copyright 2017 HootSuite Media Inc.
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Modified hereafter by contributors to runatlantis/atlantis.

package vcs

import (
	"fmt"
	"net/url"

	"github.com/cloudposse/atlantis/server/events/models"
	"github.com/lkysow/go-gitlab"
)

type GitlabClient struct {
	Client *gitlab.Client
}

// GetModifiedFiles returns the names of files that were modified in the merge request.
// The names include the path to the file from the repo root, ex. parent/child/file.txt.
func (g *GitlabClient) GetModifiedFiles(repo models.Repo, pull models.PullRequest) ([]string, error) {
	const maxPerPage = 100
	var files []string
	nextPage := 1
	// Constructing the api url by hand so we can do pagination.
	apiURL := fmt.Sprintf("projects/%s/merge_requests/%d/changes", url.QueryEscape(repo.FullName), pull.Num)
	for {
		opts := gitlab.ListOptions{
			Page:    nextPage,
			PerPage: maxPerPage,
		}
		req, err := g.Client.NewRequest("GET", apiURL, opts, nil)
		if err != nil {
			return nil, err
		}
		mr := new(gitlab.MergeRequest)
		resp, err := g.Client.Do(req, mr)
		if err != nil {
			return nil, err
		}

		for _, f := range mr.Changes {
			files = append(files, f.NewPath)
		}
		if resp.NextPage == 0 {
			break
		}
		nextPage = resp.NextPage
	}

	return files, nil
}

// CreateComment creates a comment on the merge request.
func (g *GitlabClient) CreateComment(repo models.Repo, pullNum int, comment string) error {
	_, _, err := g.Client.Notes.CreateMergeRequestNote(repo.FullName, pullNum, &gitlab.CreateMergeRequestNoteOptions{Body: gitlab.String(comment)})
	return err
}

// PullIsApproved returns true if the merge request was approved.
func (g *GitlabClient) PullIsApproved(repo models.Repo, pull models.PullRequest) (bool, error) {
	approvals, _, err := g.Client.MergeRequests.GetMergeRequestApprovals(repo.FullName, pull.Num)
	if err != nil {
		return false, err
	}
	if approvals.ApprovalsLeft > 0 {
		return false, nil
	}
	return true, nil
}

// PullIsMergeable returns true if the merge request can be merged.
// In GitLab, there isn't a single field that tells us if the pull request is
// mergeable so for now we check the merge_status and approvals_before_merge
// fields. We aren't checking if there are unresolved discussions or failing
// pipelines because those only block merges if the repo is set to require that.
// In order to check if the repo required these, we'd need to make another API
// call to get the repo settings. For now I'm going to leave this as is and if
// some users require checking this as well then we can revisit.
// It's also possible that GitLab implements their own "mergeable" field in
// their API in the future.
// See:
// - https://gitlab.com/gitlab-org/gitlab-ee/issues/3169
// - https://gitlab.com/gitlab-org/gitlab-ce/issues/42344
func (g *GitlabClient) PullIsMergeable(repo models.Repo, pull models.PullRequest) (bool, error) {
	mr, _, err := g.Client.MergeRequests.GetMergeRequest(repo.FullName, pull.Num)
	if err != nil {
		return false, err
	}
	if mr.MergeStatus == "can_be_merged" && mr.ApprovalsBeforeMerge <= 0 {
		return true, nil
	}
	return false, nil
}

// UpdateStatus updates the build status of a commit.
func (g *GitlabClient) UpdateStatus(repo models.Repo, pull models.PullRequest, state models.CommitStatus, description string) error {
	const statusContext = "Atlantis"

	gitlabState := gitlab.Failed
	switch state {
	case models.PendingCommitStatus:
		gitlabState = gitlab.Pending
	case models.FailedCommitStatus:
		gitlabState = gitlab.Failed
	case models.SuccessCommitStatus:
		gitlabState = gitlab.Success
	}
	_, _, err := g.Client.Commits.SetCommitStatus(repo.FullName, pull.HeadCommit, &gitlab.SetCommitStatusOptions{
		State:       gitlabState,
		Context:     gitlab.String(statusContext),
		Description: gitlab.String(description),
	})
	return err
}

func (g *GitlabClient) GetMergeRequest(repoFullName string, pullNum int) (*gitlab.MergeRequest, error) {
	mr, _, err := g.Client.MergeRequests.GetMergeRequest(repoFullName, pullNum)
	return mr, err
}

// GetTeamNamesForUser returns the names of the teams or groups that the user belongs to (in the organization the repository belongs to).
func (g *GitlabClient) GetTeamNamesForUser(repo models.Repo, user models.User) ([]string, error) {
	return nil, nil
}
