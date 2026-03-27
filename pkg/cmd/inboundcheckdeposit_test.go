// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestInboundCheckDepositsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-check-deposits", "retrieve",
			"--inbound-check-deposit-id", "inbound_check_deposit_zoshvqybq0cjjm31mra",
		)
	})
}

func TestInboundCheckDepositsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-check-deposits", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--check-transfer-id", "check_transfer_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(inboundCheckDepositsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-check-deposits", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--check-transfer-id", "check_transfer_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}

func TestInboundCheckDepositsDecline(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-check-deposits", "decline",
			"--inbound-check-deposit-id", "inbound_check_deposit_zoshvqybq0cjjm31mra",
		)
	})
}

func TestInboundCheckDepositsReturn(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-check-deposits", "return",
			"--inbound-check-deposit-id", "inbound_check_deposit_zoshvqybq0cjjm31mra",
			"--reason", "altered_or_fictitious",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: altered_or_fictitious")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"inbound-check-deposits", "return",
			"--inbound-check-deposit-id", "inbound_check_deposit_zoshvqybq0cjjm31mra",
		)
	})
}
