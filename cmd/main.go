package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	tferun "github.com/kvrhdn/go-tfe-run"
)

type options struct {
	// TODO we shouldn't pass the token as command line argument
	Token             string `short:"t" long:"token" required:"true" description:"Terraform Cloud token"`
	Organization      string `short:"o" long:"organization" required:"true" description:"The organization on Terraform Cloud."`
	Workspace         string `short:"w" long:"workspace" required:"true" description:"The workspace on Terraform Cloud."`
	Message           string `short:"m" long:"message" default:"Created using go-tfe-run" description:"Message to use as name of the run."`
	Directory         string `short:"d" long:"directory" default:"./" description:"The directory that is uploaded to Terraform Cloud, respects .terraformignore."`
	Speculative       bool   `long:"speculative" description:"Whether to create a speculative run. A speculative run can not be applied."`
	WaitForCompletion bool   `long:"wait" description:"Whether we should wait for the non-speculative run to be applied. This will block until the run is finished."`
	TfVars            string `long:"tf-vars" description:"Contents of a auto.tfvars file that will be uploaded to Terraform Cloud. This can be used to set Terraform variables. These variables will not be persisted."`
}

func main() {
	var opts options

	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	ctx := context.Background()

	runOptions := tferun.RunOptions{
		Token:             opts.Token,
		Organization:      opts.Organization,
		Workspace:         opts.Workspace,
		Message:           opts.Message,
		Directory:         opts.Directory,
		Speculative:       opts.Speculative,
		WaitForCompletion: opts.WaitForCompletion,
		TfVars:            opts.TfVars,
	}

	_, err = tferun.Run(ctx, runOptions)

	if err != nil {
		fmt.Printf("Error: run failed: %v", err)
		os.Exit(1)
	}
}
