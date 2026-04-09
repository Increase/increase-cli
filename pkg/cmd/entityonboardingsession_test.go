// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestEntityOnboardingSessionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-onboarding-sessions", "create",
			"--program-id", "program_i2v2os4mwza1oetokh9i",
			"--redirect-url", "https://example.com/onboarding/session",
			"--entity-id", "entity_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"program_id: program_i2v2os4mwza1oetokh9i\n" +
			"redirect_url: https://example.com/onboarding/session\n" +
			"entity_id: entity_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entity-onboarding-sessions", "create",
		)
	})
}

func TestEntityOnboardingSessionsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-onboarding-sessions", "retrieve",
			"--entity-onboarding-session-id", "entity_onboarding_session_wid2ug11fsmvh3k9hymd",
		)
	})
}

func TestEntityOnboardingSessionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-onboarding-sessions", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [active]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(entityOnboardingSessionsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-onboarding-sessions", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[active]",
		)
	})
}

func TestEntityOnboardingSessionsExpire(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-onboarding-sessions", "expire",
			"--entity-onboarding-session-id", "entity_onboarding_session_wid2ug11fsmvh3k9hymd",
		)
	})
}
