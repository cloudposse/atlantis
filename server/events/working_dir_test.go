package events_test

import (
	"fmt"
	"testing"

	"github.com/cloudposse/atlantis/server/events"
	. "github.com/cloudposse/atlantis/testing"
)

func TestCreateGitCloneCommand(t *testing.T) {
	gitVersion, err := events.GetGitVersion()
	Ok(t, err)
	fmt.Printf("Found Git version: %s\n", gitVersion.String())

	cloneCmd, err := events.CreateGitCloneCommand(gitVersion, "https://github.com/runatlantis/atlantis", "/tmp/runatlantis/atlantis")
	Ok(t, err)

	Assert(t, cloneCmd != nil, "could not create 'git clone' command")
	fmt.Printf("git clone arguments: %s\n", cloneCmd.Args)
}
