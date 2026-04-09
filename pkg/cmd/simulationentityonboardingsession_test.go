// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsEntityOnboardingSessionsSubmit(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:entity-onboarding-sessions", "submit",
			"--entity-onboarding-session-id", "entity_onboarding_session_wid2ug11fsmvh3k9hymd",
		)
	})
}
