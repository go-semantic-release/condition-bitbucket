# :white_check_mark: condition-bitbucket

[![CI](https://github.com/go-semantic-release/condition-bitbucket/workflows/CI/badge.svg?branch=main)](https://github.com/go-semantic-release/condition-bitbucket/actions?query=workflow%3ACI+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-semantic-release/condition-bitbucket)](https://goreportcard.com/report/github.com/go-semantic-release/condition-bitbucket)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-semantic-release/condition-bitbucket)](https://pkg.go.dev/github.com/go-semantic-release/condition-bitbucket)

The Bitbucket Pipelines condition for
[go-semantic-release](https://github.com/go-semantic-release/semantic-release).

This plugin will make sure that the continuous integration environment
conditions are met. Namely that the release is happening within
Bitbucket Pipelines and is on the correct branch.

## Usage

To use this plugin you need to include the following block in your
`.semrelrc` file.

```json
{
    "plugins": {
        "ci-condition": {
            "name": "bitbucket@^1.0.0",
            // Options can be omitted if you want to use the defaults.
            // See the section on configuration below.
            "options": {
                // Put configuration options here
            }
        }
    }
}
```

### Configuration

|      Name      | Default Value |                 Description                 |
|:--------------:|:-------------:|:-------------------------------------------:|
| defaultBranch  | master        | The branch where deployments should happen. |

## Licence

The [MIT License (MIT)](http://opensource.org/licenses/MIT)

Copyright © 2021 [James Durand](https://github.com/durandj)
