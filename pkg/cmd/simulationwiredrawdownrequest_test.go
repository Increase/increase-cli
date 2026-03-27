// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsWireDrawdownRequestsRefuse(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:wire-drawdown-requests", "refuse",
			"--wire-drawdown-request-id", "wire_drawdown_request_q6lmocus3glo0lr2bfv3",
		)
	})
}

func TestSimulationsWireDrawdownRequestsSubmit(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:wire-drawdown-requests", "submit",
			"--wire-drawdown-request-id", "wire_drawdown_request_q6lmocus3glo0lr2bfv3",
		)
	})
}
