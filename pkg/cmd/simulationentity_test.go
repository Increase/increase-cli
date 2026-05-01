// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSimulationsEntitiesValidation(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:entities", "validation",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--issue", "{category: entity_tax_identifier}",
			"--status", "invalid",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simulationsEntitiesValidation)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:entities", "validation",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--issue.category", "entity_tax_identifier",
			"--status", "invalid",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"issues:\n" +
			"  - category: entity_tax_identifier\n" +
			"status: invalid\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:entities", "validation",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
		)
	})
}
