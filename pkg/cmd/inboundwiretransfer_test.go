// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestInboundWireTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-wire-transfers", "retrieve",
			"--inbound-wire-transfer-id", "inbound_wire_transfer_f228m6bmhtcxjco9pwp0",
		)
	})
}

func TestInboundWireTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-wire-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--account-number-id", "account_number_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--limit", "1",
			"--status", "{in: [pending]}",
			"--wire-drawdown-request-id", "wire_drawdown_request_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(inboundWireTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-wire-transfers", "list",
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
			"--wire-drawdown-request-id", "wire_drawdown_request_id",
		)
	})
}

func TestInboundWireTransfersReverse(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-wire-transfers", "reverse",
			"--inbound-wire-transfer-id", "inbound_wire_transfer_f228m6bmhtcxjco9pwp0",
			"--reason", "creditor_request",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: creditor_request")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"inbound-wire-transfers", "reverse",
			"--inbound-wire-transfer-id", "inbound_wire_transfer_f228m6bmhtcxjco9pwp0",
		)
	})
}
