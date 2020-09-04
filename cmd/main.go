package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	tferun "github.com/kvrhdn/go-tfe-run"
)

type options struct {
	Token             string  `short:"t" long:"token" required:"true" description:"Terraform Cloud token"`
	Organization      string  `short:"o" long:"organization" required:"true" description:"The organization on Terraform Cloud."`
	Workspace         string  `short:"w" long:"workspace" required:"true" description:"The workspace on Terraform Cloud."`
	Message           *string `short:"m" long:"message" description:"Message to use as name of the run."`
	Directory         *string `short:"d" long:"directory" default:"./" description:"The directory that is uploaded to Terraform Cloud, respects .terraformignore."`
	Speculative       bool    `long:"speculative" description:"Whether to create a speculative run. A speculative run can not be applied."`
	WaitForCompletion bool    `long:"wait" description:"Whether we should wait for the non-speculative run to be applied. This will block until the run is finished."`
	TfVars            *string `long:"tf-vars" description:"Contents of a auto.tfvars file that will be uploaded to Terraform Cloud. This can be used to set Terraform variables. These variables will not be persisted."`
}

func main() {
	var opts options

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	ctx := context.Background()

	cfg := tferun.ClientConfig{
		Token:        opts.Token,
		Organization: opts.Organization,
		Workspace:    opts.Workspace,
	}
	c, err := tferun.NewClient(ctx, cfg)
	if err != nil {
		exitWithError(err)
	}

	runOptions := tferun.RunOptions{
		Message:           opts.Message,
		Directory:         opts.Directory,
		Speculative:       opts.Speculative,
		WaitForCompletion: opts.WaitForCompletion,
		TfVars:            opts.TfVars,
	}
	_, err = c.Run(ctx, runOptions)
	if err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}
