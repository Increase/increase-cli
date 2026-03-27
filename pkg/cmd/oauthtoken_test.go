// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestOAuthTokensCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth-tokens", "create",
			"--grant-type", "authorization_code",
			"--client-id", "12345",
			"--client-secret", "supersecret",
			"--code", "123",
			"--production-token", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"grant_type: authorization_code\n" +
			"client_id: '12345'\n" +
			"client_secret: supersecret\n" +
			"code: '123'\n" +
			"production_token: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"oauth-tokens", "create",
		)
	})
}
