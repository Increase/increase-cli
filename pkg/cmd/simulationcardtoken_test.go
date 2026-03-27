// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSimulationsCardTokensCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-tokens", "create",
			"--capability", "{cross_border_push_transfers: supported, domestic_push_transfers: supported, route: visa}",
			"--expiration", "'2019-12-27'",
			"--last4", "1234",
			"--outcome", "{result: approve, decline: {reason: do_not_honor}}",
			"--prefix", "41234567",
			"--primary-account-number-length", "16",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simulationsCardTokensCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-tokens", "create",
			"--capability.cross-border-push-transfers", "supported",
			"--capability.domestic-push-transfers", "supported",
			"--capability.route", "visa",
			"--expiration", "'2019-12-27'",
			"--last4", "1234",
			"--outcome.result", "approve",
			"--outcome.decline", "{reason: do_not_honor}",
			"--prefix", "41234567",
			"--primary-account-number-length", "16",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"capabilities:\n" +
			"  - cross_border_push_transfers: supported\n" +
			"    domestic_push_transfers: supported\n" +
			"    route: visa\n" +
			"expiration: '2019-12-27'\n" +
			"last4: '1234'\n" +
			"outcome:\n" +
			"  result: approve\n" +
			"  decline:\n" +
			"    reason: do_not_honor\n" +
			"prefix: '41234567'\n" +
			"primary_account_number_length: 16\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-tokens", "create",
		)
	})
}
