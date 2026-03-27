// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsInboundMailItemsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:inbound-mail-items", "create",
			"--amount", "1000",
			"--lockbox-id", "lockbox_3xt21ok13q19advds4t5",
			"--contents-file-id", "contents_file_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 1000\n" +
			"lockbox_id: lockbox_3xt21ok13q19advds4t5\n" +
			"contents_file_id: contents_file_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:inbound-mail-items", "create",
		)
	})
}
