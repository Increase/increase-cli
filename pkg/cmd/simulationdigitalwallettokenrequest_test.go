// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsDigitalWalletTokenRequestsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:digital-wallet-token-requests", "create",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("card_id: card_oubs0hwk5rn6knuecxg2")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:digital-wallet-token-requests", "create",
		)
	})
}
