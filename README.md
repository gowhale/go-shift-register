# go-shift-register

This repo is used to drive shift registers in golang. I am using this model of shift register SN74HC595. [Datasheet](https://www.ti.com/lit/ds/symlink/sn74hc595.pdf?ts=1670067468436&ref_url=https%253A%252F%252Fwww.ti.com%252Fproduct%252FSN74HC595)

## Actions created by this template:


### Testing

The pkg-cov workflow runs all go tests and ensures pkg coverage is above 80%.

![example event parameter](https://github.com/gowhale/go-shift-register/actions/workflows/pkg-cov.yml/badge.svg?event=push)

The pages workflow publishes a test coverage website everytime there is a push to the main branch. The website can be found here: https://gowhale.github.io/go-shift-register/#file0

![example event parameter](https://github.com/gowhale/go-shift-register/actions/workflows/pages.yml/badge.svg?event=push)

### Linters

The revive workflow is executed to statically analsye go files: https://github.com/mgechev/revive

![example event parameter](https://github.com/gowhale/go-shift-register/actions/workflows/revive.yml/badge.svg?event=push)

The golangci-lint workflow runs the golangci-lint linter: https://github.com/golangci/golangci-lint

![example event parameter](https://github.com/gowhale/go-shift-register/actions/workflows/golangci-lint.yml/badge.svg?event=push)

### Project Management

The issue workflow adds a new issue to the projects Kanban board:

![example event parameter](https://github.com/gowhale/go-shift-register/actions/workflows/issue.yml/badge.svg?event=push)

The cut release workflow creates a binary executable everytime a release is published. The binary file is attached to the release.

![example event parameter](https://github.com/gowhale/go-shift-register/actions/workflows/cut-release.yml/badge.svg?event=push)

