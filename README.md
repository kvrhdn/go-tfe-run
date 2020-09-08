# `go-tfe-run` [![PkgGoDev](https://pkg.go.dev/badge/github.com/kvrhdn/go-tfe-run)](https://pkg.go.dev/github.com/kvrhdn/go-tfe-run?tab=doc)

[![CI](https://github.com/kvrhdn/go-tfe-run/workflows/CI/badge.svg)](https://github.com/kvrhdn/go-tfe-run/actions?query=workflow%3ACI)
[![Integration](https://github.com/kvrhdn/go-tfe-run/workflows/Integration/badge.svg)](https://github.com/kvrhdn/go-tfe-run/actions?query=workflow%3AIntegration)
[![Go Report Card](https://goreportcard.com/badge/github.com/kvrhdn/go-tfe-run)](https://goreportcard.com/report/github.com/kvrhdn/go-tfe-run)

A Go library to create and follow up on a run on Terraform Cloud. This library uses the [Terraform Cloud API](https://www.terraform.io/docs/cloud/run/api.html) and does not rely on a local Terraform installation.

When is ths useful?

- You want to configure parameters that are not available using the CLI, for example the name of the run or whether it is a speculative plan.
- You want to schedule a run but not wait for its completion.
- You are in an environment without Terraform CLI installed.

If you wish to integrate this into your GitHub Actions workflows, checkout the [tfe-run action](https://github.com/marketplace/actions/tfe-run) which wraps `go-tfe-run` into a custom action.

Visit [the documentation on pkg.go.dev](https://pkg.go.dev/github.com/kvrhdn/go-tfe-run?tab=doc).

## License

This software is distributed under the terms of the MIT license, see [LICENSE](./LICENSE) for details.
