// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestInboundWireDrawdownRequestsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-wire-drawdown-requests", "retrieve",
			"--inbound-wire-drawdown-request-id", "inbound_wire_drawdown_request_u5a92ikqhz1ytphn799e",
		)
	})
}

func TestInboundWireDrawdownRequestsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-wire-drawdown-requests", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}
