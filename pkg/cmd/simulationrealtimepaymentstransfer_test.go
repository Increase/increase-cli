// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSimulationsRealTimePaymentsTransfersComplete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:real-time-payments-transfers", "complete",
			"--real-time-payments-transfer-id", "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
			"--rejection", "{reject_reason_code: account_closed}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simulationsRealTimePaymentsTransfersComplete)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:real-time-payments-transfers", "complete",
			"--real-time-payments-transfer-id", "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
			"--rejection.reject-reason-code", "account_closed",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"rejection:\n" +
			"  reject_reason_code: account_closed\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:real-time-payments-transfers", "complete",
			"--real-time-payments-transfer-id", "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
		)
	})
}
