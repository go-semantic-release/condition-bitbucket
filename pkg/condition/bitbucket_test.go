package condition

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiresBranchBuildOnly(t *testing.T) {
	bbp := BitbucketPipelines{}

	setTestBranch("")

	err := bbp.RunCondition(createTestPluginConfig())

	assert.EqualError(t, err, "This test run is not running on a branch build")
}

func TestRequiresCurrentBranchToMatchTargetBranch(t *testing.T) {
	bbp := BitbucketPipelines{}

	setTestBranch("test")

	err := bbp.RunCondition(createTestPluginConfig())

	assert.EqualError(
		t,
		err,
		"This test run was triggered on the branch test, while semantic-release is configured to only publish from main",
	)
}

func TestRunsWhenAllConditionsAreMet(t *testing.T) {
	bbp := BitbucketPipelines{}

	setTestBranch("main")

	err := bbp.RunCondition(createTestPluginConfig())

	assert.Nil(t, err)
}

func setTestBranch(branch string) {
	_ = os.Setenv(bitbucketBranch, branch)
}

func createTestPluginConfig() map[string]string {
	return map[string]string{
		"defaultBranch": "main",
	}
}
