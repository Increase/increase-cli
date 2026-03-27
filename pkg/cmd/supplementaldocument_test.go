// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSupplementalDocumentsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"supplemental-documents", "create",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--file-id", "file_makxrc67oh9l6sg7w9yc",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entity_id: entity_n8y8tnk2p9339ti393yi\n" +
			"file_id: file_makxrc67oh9l6sg7w9yc\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"supplemental-documents", "create",
		)
	})
}

func TestSupplementalDocumentsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"supplemental-documents", "list",
			"--max-items", "10",
			"--entity-id", "entity_id",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})
}
