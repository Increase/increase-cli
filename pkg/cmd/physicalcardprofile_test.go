// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestPhysicalCardProfilesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-card-profiles", "create",
			"--carrier-image-file-id", "file_h6v7mtipe119os47ehlu",
			"--contact-phone", "+16505046304",
			"--description", "My Card Profile",
			"--front-image-file-id", "file_o6aex13wm1jcc36sgcj1",
			"--program-id", "program_i2v2os4mwza1oetokh9i",
			"--front-text", "{line1: x, line2: x}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(physicalCardProfilesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-card-profiles", "create",
			"--carrier-image-file-id", "file_h6v7mtipe119os47ehlu",
			"--contact-phone", "+16505046304",
			"--description", "My Card Profile",
			"--front-image-file-id", "file_o6aex13wm1jcc36sgcj1",
			"--program-id", "program_i2v2os4mwza1oetokh9i",
			"--front-text.line1", "x",
			"--front-text.line2", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"carrier_image_file_id: file_h6v7mtipe119os47ehlu\n" +
			"contact_phone: '+16505046304'\n" +
			"description: My Card Profile\n" +
			"front_image_file_id: file_o6aex13wm1jcc36sgcj1\n" +
			"program_id: program_i2v2os4mwza1oetokh9i\n" +
			"front_text:\n" +
			"  line1: x\n" +
			"  line2: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"physical-card-profiles", "create",
		)
	})
}

func TestPhysicalCardProfilesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-card-profiles", "retrieve",
			"--physical-card-profile-id", "physical_card_profile_m534d5rn9qyy9ufqxoec",
		)
	})
}

func TestPhysicalCardProfilesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-card-profiles", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending_creating]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(physicalCardProfilesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-card-profiles", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[pending_creating]",
		)
	})
}

func TestPhysicalCardProfilesArchive(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-card-profiles", "archive",
			"--physical-card-profile-id", "physical_card_profile_m534d5rn9qyy9ufqxoec",
		)
	})
}

func TestPhysicalCardProfilesClone(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-card-profiles", "clone",
			"--physical-card-profile-id", "physical_card_profile_m534d5rn9qyy9ufqxoec",
			"--carrier-image-file-id", "carrier_image_file_id",
			"--contact-phone", "x",
			"--description", "x",
			"--front-image-file-id", "file_o6aex13wm1jcc36sgcj1",
			"--front-text", "{line1: x, line2: x}",
			"--program-id", "program_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(physicalCardProfilesClone)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-card-profiles", "clone",
			"--physical-card-profile-id", "physical_card_profile_m534d5rn9qyy9ufqxoec",
			"--carrier-image-file-id", "carrier_image_file_id",
			"--contact-phone", "x",
			"--description", "x",
			"--front-image-file-id", "file_o6aex13wm1jcc36sgcj1",
			"--front-text.line1", "x",
			"--front-text.line2", "x",
			"--program-id", "program_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"carrier_image_file_id: carrier_image_file_id\n" +
			"contact_phone: x\n" +
			"description: x\n" +
			"front_image_file_id: file_o6aex13wm1jcc36sgcj1\n" +
			"front_text:\n" +
			"  line1: x\n" +
			"  line2: x\n" +
			"program_id: program_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"physical-card-profiles", "clone",
			"--physical-card-profile-id", "physical_card_profile_m534d5rn9qyy9ufqxoec",
		)
	})
}
