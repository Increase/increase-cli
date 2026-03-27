// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestInboundACHTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-ach-transfers", "retrieve",
			"--inbound-ach-transfer-id", "inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
		)
	})
}

func TestInboundACHTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-ach-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--account-number-id", "account_number_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--limit", "1",
			"--status", "{in: [pending]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(inboundACHTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-ach-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--account-number-id", "account_number_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--limit", "1",
			"--status.in", "[pending]",
		)
	})
}

func TestInboundACHTransfersCreateNotificationOfChange(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-ach-transfers", "create-notification-of-change",
			"--inbound-ach-transfer-id", "inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
			"--updated-account-number", "987654321",
			"--updated-routing-number", "101050001",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"updated_account_number: '987654321'\n" +
			"updated_routing_number: '101050001'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"inbound-ach-transfers", "create-notification-of-change",
			"--inbound-ach-transfer-id", "inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
		)
	})
}

func TestInboundACHTransfersDecline(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-ach-transfers", "decline",
			"--inbound-ach-transfer-id", "inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
			"--reason", "payment_stopped",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: payment_stopped")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"inbound-ach-transfers", "decline",
			"--inbound-ach-transfer-id", "inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
		)
	})
}

func TestInboundACHTransfersTransferReturn(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-ach-transfers", "transfer-return",
			"--inbound-ach-transfer-id", "inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
			"--reason", "payment_stopped",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: payment_stopped")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"inbound-ach-transfers", "transfer-return",
			"--inbound-ach-transfer-id", "inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
		)
	})
}
