// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestRoutingNumbersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-numbers", "list",
			"--max-items", "10",
			"--routing-number", "483310694",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}
