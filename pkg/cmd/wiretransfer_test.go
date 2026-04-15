// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestWireTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "100",
			"--creditor", "{name: Ian Crease, address: {unstructured: {line1: 33 Liberty Street, line2: New York, line3: NY 10045}}}",
			"--remittance", "{category: unstructured, tax: {date: '2019-12-27', identification_number: '483310694', type_code: 1I5r3}, unstructured: {message: New account transfer}}",
			"--account-number", "987654321",
			"--debtor", "{name: name, address: {unstructured: {line1: line1, line2: line2, line3: line3}}}",
			"--external-account-id", "external_account_id",
			"--inbound-wire-drawdown-request-id", "inbound_wire_drawdown_request_id",
			"--require-approval=true",
			"--routing-number", "101050001",
			"--source-account-number-id", "source_account_number_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(wireTransfersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "100",
			"--creditor.name", "Ian Crease",
			"--creditor.address", "{unstructured: {line1: 33 Liberty Street, line2: New York, line3: NY 10045}}",
			"--remittance.category", "unstructured",
			"--remittance.tax", "{date: '2019-12-27', identification_number: '483310694', type_code: 1I5r3}",
			"--remittance.unstructured", "{message: New account transfer}",
			"--account-number", "987654321",
			"--debtor.name", "name",
			"--debtor.address", "{unstructured: {line1: line1, line2: line2, line3: line3}}",
			"--external-account-id", "external_account_id",
			"--inbound-wire-drawdown-request-id", "inbound_wire_drawdown_request_id",
			"--require-approval=true",
			"--routing-number", "101050001",
			"--source-account-number-id", "source_account_number_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"amount: 100\n" +
			"creditor:\n" +
			"  name: Ian Crease\n" +
			"  address:\n" +
			"    unstructured:\n" +
			"      line1: 33 Liberty Street\n" +
			"      line2: New York\n" +
			"      line3: NY 10045\n" +
			"remittance:\n" +
			"  category: unstructured\n" +
			"  tax:\n" +
			"    date: '2019-12-27'\n" +
			"    identification_number: '483310694'\n" +
			"    type_code: 1I5r3\n" +
			"  unstructured:\n" +
			"    message: New account transfer\n" +
			"account_number: '987654321'\n" +
			"debtor:\n" +
			"  name: name\n" +
			"  address:\n" +
			"    unstructured:\n" +
			"      line1: line1\n" +
			"      line2: line2\n" +
			"      line3: line3\n" +
			"external_account_id: external_account_id\n" +
			"inbound_wire_drawdown_request_id: inbound_wire_drawdown_request_id\n" +
			"require_approval: true\n" +
			"routing_number: '101050001'\n" +
			"source_account_number_id: source_account_number_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"wire-transfers", "create",
		)
	})
}

func TestWireTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-transfers", "retrieve",
			"--wire-transfer-id", "wire_transfer_5akynk7dqsq25qwk9q2u",
		)
	})
}

func TestWireTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--external-account-id", "external_account_id",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending_approval]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(wireTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--external-account-id", "external_account_id",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[pending_approval]",
		)
	})
}

func TestWireTransfersApprove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-transfers", "approve",
			"--wire-transfer-id", "wire_transfer_5akynk7dqsq25qwk9q2u",
		)
	})
}

func TestWireTransfersCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-transfers", "cancel",
			"--wire-transfer-id", "wire_transfer_5akynk7dqsq25qwk9q2u",
		)
	})
}
