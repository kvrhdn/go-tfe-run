package tferun

import (
	"context"
	"fmt"
	"os"
)

func Example() {
	token := os.Getenv("TFE_TOKEN")

	options := RunOptions{
		Token:        token,
		Organization: "kvrhdn",
		Workspace:    "tfe-run",
		Message:      String("Run created using tfe-run"),
	}
	output, err := Run(context.TODO(), options)
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
	}

	fmt.Printf("Run created: %v", output.RunURL)
}

func Example_tfVars() {
	token := os.Getenv("TFE_TOKEN")

	options := RunOptions{
		Token:        token,
		Organization: "kvrhdn",
		Workspace:    "tfe-run",
		TfVars:       String("my_var = \"foo\""),
	}
	output, err := Run(context.TODO(), options)
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
	}

	fmt.Printf("Run created: %v", output.RunURL)
}
