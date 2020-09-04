package tferun

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	ctx := context.Background()

	if testing.Short() {
		t.Skip()
	}

	token, ok := os.LookupEnv("TFE_TOKEN")
	if !ok {
		t.Fatal("Expected environment variable TFE_TOKEN to be set")
	}
	runNr, ok := os.LookupEnv("GITHUB_RUN_NUMBER")
	if !ok {
		t.Fatal("Expected environment variable GITHUB_RUN_NUMBER to be set")
	}

	cfg := ClientConfig{
		Token:        token,
		Organization: "kvrhdn",
		Workspace:    "go-tfe-run",
	}
	client, err := NewClient(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Speculative run, with changes", func(t *testing.T) {
		planOptions := RunOptions{
			Message:           String(fmt.Sprintf("Plan for %s", runNr)),
			Directory:         String("./testdata"),
			Speculative:       true,
			WaitForCompletion: true,
			TfVars:            String(fmt.Sprintf("github_run_number = \"%s\"", runNr)),
		}
		planOutput, err := client.Run(ctx, planOptions)
		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, planOutput.RunURL, "https://app.terraform.io/app/kvrhdn/workspaces/go-tfe-run/runs/run-")
		assert.Equal(t, true, *planOutput.HasChanges)
	})

	t.Run("Non-speculative run, with changes", func(t *testing.T) {
		applyOptions := RunOptions{
			Message:           String(fmt.Sprintf("Apply of run %s", runNr)),
			Directory:         String("./testdata"),
			Speculative:       false,
			WaitForCompletion: true,
			TfVars:            String(fmt.Sprintf("github_run_number = \"%s\"", runNr)),
		}
		applyOutput, err := client.Run(ctx, applyOptions)
		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, applyOutput.RunURL, "https://app.terraform.io/app/kvrhdn/workspaces/go-tfe-run/runs/run-")
		assert.Equal(t, true, *applyOutput.HasChanges)

		expectedOutputs := map[string]string{
			"marker_message": fmt.Sprintf("Integration - run %s", runNr),
		}
		assert.Equal(t, expectedOutputs, *applyOutput.TfOutputs)
	})

	t.Run("Speculative run, no changes", func(t *testing.T) {
		planOptions := RunOptions{
			Message:           String(fmt.Sprintf("Plan for %s", runNr)),
			Directory:         String("./testdata"),
			Speculative:       true,
			WaitForCompletion: true,
			TfVars:            String(fmt.Sprintf("github_run_number = \"%s\"", runNr)),
		}
		planOutput, err := client.Run(ctx, planOptions)
		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, planOutput.RunURL, "https://app.terraform.io/app/kvrhdn/workspaces/go-tfe-run/runs/run-")
		assert.Equal(t, false, *planOutput.HasChanges)
	})
}
