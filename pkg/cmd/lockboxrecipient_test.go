// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestLockboxRecipientsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lockbox-recipients", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--lockbox-address-id", "lockbox_address_lw6sbzl9ol5dfd8hdml6",
			"--description", "x",
			"--recipient-name", "Ian Crease",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"lockbox_address_id: lockbox_address_lw6sbzl9ol5dfd8hdml6\n" +
			"description: x\n" +
			"recipient_name: Ian Crease\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"lockbox-recipients", "create",
		)
	})
}

func TestLockboxRecipientsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lockbox-recipients", "retrieve",
			"--lockbox-recipient-id", "lockbox_3xt21ok13q19advds4t5",
		)
	})
}

func TestLockboxRecipientsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lockbox-recipients", "update",
			"--lockbox-recipient-id", "lockbox_3xt21ok13q19advds4t5",
			"--description", "x",
			"--recipient-name", "x",
			"--status", "active",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: x\n" +
			"recipient_name: x\n" +
			"status: active\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"lockbox-recipients", "update",
			"--lockbox-recipient-id", "lockbox_3xt21ok13q19advds4t5",
		)
	})
}

func TestLockboxRecipientsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lockbox-recipients", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--lockbox-address-id", "lockbox_address_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(lockboxRecipientsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lockbox-recipients", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--lockbox-address-id", "lockbox_address_id",
		)
	})
}
