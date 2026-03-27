// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestIntrafiExclusionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-exclusions", "create",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--fdic-certificate-number", "314159",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entity_id: entity_n8y8tnk2p9339ti393yi\n" +
			"fdic_certificate_number: '314159'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"intrafi-exclusions", "create",
		)
	})
}

func TestIntrafiExclusionsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-exclusions", "retrieve",
			"--intrafi-exclusion-id", "account_in71c4amph0vgo2qllky",
		)
	})
}

func TestIntrafiExclusionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-exclusions", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--entity-id", "entity_id",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})
}

func TestIntrafiExclusionsArchive(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-exclusions", "archive",
			"--intrafi-exclusion-id", "intrafi_exclusion_ygfqduuzpau3jqof6jyh",
		)
	})
}
