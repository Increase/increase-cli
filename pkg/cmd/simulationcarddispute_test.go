// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSimulationsCardDisputesAction(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-disputes", "action",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
			"--network", "visa",
			"--visa", "{action: accept_user_submission, accept_chargeback: {}, accept_user_submission: {}, decline_user_prearbitration: {}, receive_merchant_prearbitration: {}, represent: {}, request_further_information: {reason: x}, time_out_chargeback: {}, time_out_merchant_prearbitration: {}, time_out_representment: {}, time_out_user_prearbitration: {}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simulationsCardDisputesAction)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-disputes", "action",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
			"--network", "visa",
			"--visa.action", "accept_user_submission",
			"--visa.accept-chargeback", "{}",
			"--visa.accept-user-submission", "{}",
			"--visa.decline-user-prearbitration", "{}",
			"--visa.receive-merchant-prearbitration", "{}",
			"--visa.represent", "{}",
			"--visa.request-further-information", "{reason: x}",
			"--visa.time-out-chargeback", "{}",
			"--visa.time-out-merchant-prearbitration", "{}",
			"--visa.time-out-representment", "{}",
			"--visa.time-out-user-prearbitration", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"network: visa\n" +
			"visa:\n" +
			"  action: accept_user_submission\n" +
			"  accept_chargeback: {}\n" +
			"  accept_user_submission: {}\n" +
			"  decline_user_prearbitration: {}\n" +
			"  receive_merchant_prearbitration: {}\n" +
			"  represent: {}\n" +
			"  request_further_information:\n" +
			"    reason: x\n" +
			"  time_out_chargeback: {}\n" +
			"  time_out_merchant_prearbitration: {}\n" +
			"  time_out_representment: {}\n" +
			"  time_out_user_prearbitration: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-disputes", "action",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
		)
	})
}
