// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsACHTransfersAcknowledge(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:ach-transfers", "acknowledge",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
		)
	})
}

func TestSimulationsACHTransfersCreateNotificationOfChange(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:ach-transfers", "create-notification-of-change",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
			"--corrected-account-funding", "checking",
			"--corrected-account-number", "x",
			"--corrected-individual-id", "x",
			"--corrected-routing-number", "123456789",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"corrected_account_funding: checking\n" +
			"corrected_account_number: x\n" +
			"corrected_individual_id: x\n" +
			"corrected_routing_number: '123456789'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:ach-transfers", "create-notification-of-change",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
		)
	})
}

func TestSimulationsACHTransfersReturn(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:ach-transfers", "return",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
			"--reason", "insufficient_fund",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: insufficient_fund")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:ach-transfers", "return",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
		)
	})
}

func TestSimulationsACHTransfersSettle(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:ach-transfers", "settle",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
			"--inbound-funds-hold-behavior", "release_immediately",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("inbound_funds_hold_behavior: release_immediately")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:ach-transfers", "settle",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
		)
	})
}

func TestSimulationsACHTransfersSubmit(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:ach-transfers", "submit",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
		)
	})
}
