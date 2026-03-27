// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestOAuthConnectionsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth-connections", "retrieve",
			"--oauth-connection-id", "connection_dauknoksyr4wilz4e6my",
		)
	})
}

func TestOAuthConnectionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth-connections", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "1",
			"--oauth-application-id", "oauth_application_id",
			"--status", "{in: [active]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(oauthConnectionsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth-connections", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "1",
			"--oauth-application-id", "oauth_application_id",
			"--status.in", "[active]",
		)
	})
}
