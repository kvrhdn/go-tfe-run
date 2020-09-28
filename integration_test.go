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

	t.Run("Plan - should have changes", func(t *testing.T) {
		planOptions := RunOptions{
			Message:           String(fmt.Sprintf("Integration run %s - plan", runNr)),
			Directory:         String("./testdata"),
			Type:              RunTypePlan,
			WaitForCompletion: true,
			TfVars:            String(fmt.Sprintf("github_run_number = \"%s\"", runNr)),
		}
		planOutput, err := client.Run(ctx, planOptions)

		assert.NoError(t, err)
		assert.Contains(t, planOutput.RunURL, "https://app.terraform.io/app/kvrhdn/workspaces/go-tfe-run/runs/run-")
		assert.Equal(t, true, *planOutput.HasChanges)
	})

	t.Run("Apply", func(t *testing.T) {
		applyOptions := RunOptions{
			Message:           String(fmt.Sprintf("Integration run %s - apply", runNr)),
			Directory:         String("./testdata"),
			Type:              RunTypeApply,
			WaitForCompletion: true,
			TfVars:            String(fmt.Sprintf("github_run_number = \"%s\"", runNr)),
		}
		applyOutput, err := client.Run(ctx, applyOptions)

		assert.NoError(t, err)
		assert.Contains(t, applyOutput.RunURL, "https://app.terraform.io/app/kvrhdn/workspaces/go-tfe-run/runs/run-")
		assert.Equal(t, true, *applyOutput.HasChanges)
	})

	t.Run("Get terraform outputs", func(t *testing.T) {
		outputs, err := client.GetTerraformOutputs(ctx)

		assert.NoError(t, err)

		expectedOutputs := map[string]string{
			"marker_message_1": fmt.Sprintf("Integration - run %s - marker #1", runNr),
			"marker_message_2": fmt.Sprintf("Integration - run %s - marker #2", runNr),
		}
		assert.Equal(t, expectedOutputs, outputs)
	})

	t.Run("Plan - should have no changes", func(t *testing.T) {
		planOptions := RunOptions{
			Message:           String(fmt.Sprintf("Integration run %s - plan", runNr)),
			Directory:         String("./testdata"),
			Type:              RunTypePlan,
			WaitForCompletion: true,
			TfVars:            String(fmt.Sprintf("github_run_number = \"%s\"", runNr)),
		}
		planOutput, err := client.Run(ctx, planOptions)

		assert.NoError(t, err)
		assert.Contains(t, planOutput.RunURL, "https://app.terraform.io/app/kvrhdn/workspaces/go-tfe-run/runs/run-")
		assert.Equal(t, false, *planOutput.HasChanges)
	})

	t.Run("Destroy - with address targeting", func(t *testing.T) {
		planOptions := RunOptions{
			Message:           String(fmt.Sprintf("Integration run %s - destroy", runNr)),
			Directory:         String("./testdata"),
			Type:              RunTypeDestroy,
			TargetAddrs:       []string{"honeycombio_marker.dummy_resource_1"},
			WaitForCompletion: true,
			TfVars:            String(fmt.Sprintf("github_run_number = \"%s\"", runNr)),
		}
		planOutput, err := client.Run(ctx, planOptions)

		assert.NoError(t, err)
		assert.Contains(t, planOutput.RunURL, "https://app.terraform.io/app/kvrhdn/workspaces/go-tfe-run/runs/run-")
		assert.Equal(t, true, *planOutput.HasChanges)
	})

	t.Run("Get terraform outputs - should only have marker #2", func(t *testing.T) {
		outputs, err := client.GetTerraformOutputs(ctx)

		assert.NoError(t, err)

		expectedOutputs := map[string]string{
			"marker_message_2": fmt.Sprintf("Integration - run %s - marker #2", runNr),
		}
		assert.Equal(t, expectedOutputs, outputs)
	})

	t.Run("Destroy", func(t *testing.T) {
		planOptions := RunOptions{
			Message:           String(fmt.Sprintf("Integration run %s - destroy", runNr)),
			Directory:         String("./testdata"),
			Type:              RunTypeDestroy,
			WaitForCompletion: true,
			TfVars:            String(fmt.Sprintf("github_run_number = \"%s\"", runNr)),
		}
		planOutput, err := client.Run(ctx, planOptions)

		assert.NoError(t, err)
		assert.Contains(t, planOutput.RunURL, "https://app.terraform.io/app/kvrhdn/workspaces/go-tfe-run/runs/run-")
		assert.Equal(t, true, *planOutput.HasChanges)
	})

	t.Run("Get terraform outputs - should be empty", func(t *testing.T) {
		outputs, err := client.GetTerraformOutputs(ctx)

		assert.NoError(t, err)
		assert.Empty(t, outputs)
	})
}
