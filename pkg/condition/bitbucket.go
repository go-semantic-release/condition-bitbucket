package condition

import (
	"fmt"
	"os"
)

var (
	// CIVersion is the version of the current release according to CI
	CIVersion = "dev"
)

const (
	bitbucketBranch = "BITBUCKET_BRANCH"
	bitbucketCommit = "BITBUCKET_COMMIT"
)

// BitbucketPipelines provides CI conditional handling for Bitbucket
// Pipelines.
type BitbucketPipelines struct {
}

// Name returns the name of the Semantic Release plugin.
func (bbp *BitbucketPipelines) Name() string {
	return "Bitbucket Pipelines"
}

// Version returns the version of this Semantic Release plugin.
func (bbp *BitbucketPipelines) Version() string {
	return CIVersion
}

// GetCurrentBranch returns the Git branch that Semantic Release is
// being run on.
func (bbp *BitbucketPipelines) GetCurrentBranch() string {
	return os.Getenv(bitbucketBranch)
}

// GetCurrentSHA returns the Git commit hash that Semantic Release is
// being run on.
func (bbp *BitbucketPipelines) GetCurrentSHA() string {
	return os.Getenv(bitbucketCommit)
}

// IsBranchRef determins if the current Git ref is for a branch or
// something else (Git tag, detached head, etc).
func (bbp *BitbucketPipelines) IsBranchRef() bool {
	return bbp.GetCurrentBranch() != ""
}

// RunCondition does the actual work of this plugin. It's primary
// function is to make sure that the CI environment is setup correctly
// for creating a new release.
func (bbp *BitbucketPipelines) RunCondition(config map[string]string) error {
	defaultBranch := config["defaultBranch"]

	if !bbp.IsBranchRef() {
		return fmt.Errorf("This test run is not running on a branch build")
	}

	if branch := bbp.GetCurrentBranch(); defaultBranch != "*" && branch != defaultBranch {
		return fmt.Errorf(
			"This test run was triggered on the branch %s, while semantic-release is configured to only publish from %s",
			branch,
			defaultBranch,
		)
	}

	return nil
}
