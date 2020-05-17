# `go-tfe-run`

A command line utility and Go library to create a run on Terraform Cloud. This utility will use the [Terraform Cloud API][api], and does not rely on a local Terraform installation.

[api]: https://www.terraform.io/docs/cloud/run/api.html

Why would you want to use this?

- You are in an environment without Terraform CLI.
- You want to configure parameters that are not available using the CLI: you can for example set the name of the run or create a speculative run.

## How to use this

_⚠️ Work in progress: the CLI still has some pretty rough edges!_

Install `go-tfe-run` using go get:

```
go get github.com/kvrhdn/go-tfe-run
```

Run it:

```
go-tfe-run
```

Alternatively, clone this repository and run it directly:

```
go run .
```

## Using the library

Import `go-tfe-run/lib`:

```go
import (
    tfe "github.com/kvhrnd/go-tfe-run/lib"
)
```

The library contains just one function, `Run`:

```go
options := tfe.RunOptions{
    Token:        "<your API token>",
    Organization: "kvrhdn",
    Workspace:    "go-tfe-run",
    Message:      "This run was created using go-tfe-run",
    Directory:    "./",
    Speculative:  false,
    TfVars:       "",
}

output, err := tfe.Run(ctx, options)

// handle err
```

## License

This software is distributed under the terms of the MIT license, see [LICENSE](./LICENSE) for details.
