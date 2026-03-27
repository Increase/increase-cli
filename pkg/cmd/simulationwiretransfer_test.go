// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsWireTransfersReverse(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:wire-transfers", "reverse",
			"--wire-transfer-id", "wire_transfer_5akynk7dqsq25qwk9q2u",
		)
	})
}

func TestSimulationsWireTransfersSubmit(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:wire-transfers", "submit",
			"--wire-transfer-id", "wire_transfer_5akynk7dqsq25qwk9q2u",
		)
	})
}
