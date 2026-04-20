// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestInboundMailItemsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-mail-items", "retrieve",
			"--inbound-mail-item-id", "inbound_mail_item_q6rrg7mmqpplx80zceev",
		)
	})
}

func TestInboundMailItemsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-mail-items", "list",
			"--max-items", "10",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--limit", "1",
			"--lockbox-id", "lockbox_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(inboundMailItemsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-mail-items", "list",
			"--max-items", "10",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--limit", "1",
			"--lockbox-id", "lockbox_id",
		)
	})
}

func TestInboundMailItemsAction(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-mail-items", "action",
			"--inbound-mail-item-id", "inbound_mail_item_q6rrg7mmqpplx80zceev",
			"--check", "{action: deposit, account_id: account_in71c4amph0vgo2qllky}",
			"--check", "{action: ignore, account_id: account_in71c4amph0vgo2qllky}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(inboundMailItemsAction)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound-mail-items", "action",
			"--inbound-mail-item-id", "inbound_mail_item_q6rrg7mmqpplx80zceev",
			"--check.action", "deposit",
			"--check.account-id", "account_in71c4amph0vgo2qllky",
			"--check.action", "ignore",
			"--check.account-id", "account_in71c4amph0vgo2qllky",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"checks:\n" +
			"  - action: deposit\n" +
			"    account_id: account_in71c4amph0vgo2qllky\n" +
			"  - action: ignore\n" +
			"    account_id: account_in71c4amph0vgo2qllky\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"inbound-mail-items", "action",
			"--inbound-mail-item-id", "inbound_mail_item_q6rrg7mmqpplx80zceev",
		)
	})
}
