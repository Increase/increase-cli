// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestProgramsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"programs", "retrieve",
			"--program-id", "program_i2v2os4mwza1oetokh9i",
		)
	})
}

func TestProgramsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"programs", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}
