// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSimulationsCheckDepositsAdjustment(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:check-deposits", "adjustment",
			"--check-deposit-id", "check_deposit_f06n9gpg7sxn8t19lfc1",
			"--amount", "-1000000000",
			"--reason", "late_return",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: -1000000000\n" +
			"reason: late_return\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:check-deposits", "adjustment",
			"--check-deposit-id", "check_deposit_f06n9gpg7sxn8t19lfc1",
		)
	})
}

func TestSimulationsCheckDepositsReject(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:check-deposits", "reject",
			"--check-deposit-id", "check_deposit_f06n9gpg7sxn8t19lfc1",
		)
	})
}

func TestSimulationsCheckDepositsReturn(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:check-deposits", "return",
			"--check-deposit-id", "check_deposit_f06n9gpg7sxn8t19lfc1",
		)
	})
}

func TestSimulationsCheckDepositsSubmit(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:check-deposits", "submit",
			"--check-deposit-id", "check_deposit_f06n9gpg7sxn8t19lfc1",
			"--scan", "{account_number: x, routing_number: x, auxiliary_on_us: x}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simulationsCheckDepositsSubmit)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:check-deposits", "submit",
			"--check-deposit-id", "check_deposit_f06n9gpg7sxn8t19lfc1",
			"--scan.account-number", "x",
			"--scan.routing-number", "x",
			"--scan.auxiliary-on-us", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"scan:\n" +
			"  account_number: x\n" +
			"  routing_number: x\n" +
			"  auxiliary_on_us: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:check-deposits", "submit",
			"--check-deposit-id", "check_deposit_f06n9gpg7sxn8t19lfc1",
		)
	})
}
