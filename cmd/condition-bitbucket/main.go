package main

import (
	"github.com/go-semantic-release/semantic-release/v2/pkg/condition"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"

	bitbucketCondition "github.com/durandj/semantic-release-condition-bitbucket/pkg/condition"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		CICondition: func() condition.CICondition {
			return &bitbucketCondition.BitbucketPipelines{}
		},
	})
}
