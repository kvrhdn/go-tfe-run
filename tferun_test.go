package tferun

import (
	"context"
	"fmt"
	"os"
)

func Example() {
	ctx := context.TODO()
	token := os.Getenv("TFE_TOKEN")

	cfg := ClientConfig{
		Token:        token,
		Organization: "kvrhdn",
		Workspace:    "tfe-run",
	}
	client, err := NewClient(ctx, cfg)
	if err != nil {
		fmt.Printf("could not create client: %v\n", err)
		os.Exit(1)
	}

	options := RunOptions{
		Message: String("Run created using tfe-run"),
	}
	output, err := client.Run(ctx, options)
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Run created: %v", output.RunURL)
}

func Example_tfVars() {
	ctx := context.TODO()
	token := os.Getenv("TFE_TOKEN")

	cfg := ClientConfig{
		Token:        token,
		Organization: "kvrhdn",
		Workspace:    "tfe-run",
	}
	client, err := NewClient(ctx, cfg)
	if err != nil {
		fmt.Printf("could not create client: %v\n", err)
		os.Exit(1)
	}

	options := RunOptions{
		TfVars: String("my_var = \"foo\""),
	}
	output, err := client.Run(ctx, options)
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Run created: %v", output.RunURL)
}
