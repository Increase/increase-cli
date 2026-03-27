// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestFileLinksCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"file-links", "create",
			"--file-id", "file_makxrc67oh9l6sg7w9yc",
			"--expires-at", "'2019-12-27T18:11:19.117Z'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"file_id: file_makxrc67oh9l6sg7w9yc\n" +
			"expires_at: '2019-12-27T18:11:19.117Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"file-links", "create",
		)
	})
}
