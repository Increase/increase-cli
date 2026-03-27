// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestDigitalCardProfilesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digital-card-profiles", "create",
			"--app-icon-file-id", "file_8zxqkwlh43wo144u8yec",
			"--background-image-file-id", "file_1ai913suu1zfn1pdetru",
			"--card-description", "MyBank Signature Card",
			"--description", "My Card Profile",
			"--issuer-name", "MyBank",
			"--contact-email", "user@example.com",
			"--contact-phone", "+18885551212",
			"--contact-website", "https://example.com",
			"--text-color", "{blue: 59, green: 43, red: 26}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(digitalCardProfilesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digital-card-profiles", "create",
			"--app-icon-file-id", "file_8zxqkwlh43wo144u8yec",
			"--background-image-file-id", "file_1ai913suu1zfn1pdetru",
			"--card-description", "MyBank Signature Card",
			"--description", "My Card Profile",
			"--issuer-name", "MyBank",
			"--contact-email", "user@example.com",
			"--contact-phone", "+18885551212",
			"--contact-website", "https://example.com",
			"--text-color.blue", "59",
			"--text-color.green", "43",
			"--text-color.red", "26",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"app_icon_file_id: file_8zxqkwlh43wo144u8yec\n" +
			"background_image_file_id: file_1ai913suu1zfn1pdetru\n" +
			"card_description: MyBank Signature Card\n" +
			"description: My Card Profile\n" +
			"issuer_name: MyBank\n" +
			"contact_email: user@example.com\n" +
			"contact_phone: '+18885551212'\n" +
			"contact_website: https://example.com\n" +
			"text_color:\n" +
			"  blue: 59\n" +
			"  green: 43\n" +
			"  red: 26\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"digital-card-profiles", "create",
		)
	})
}

func TestDigitalCardProfilesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digital-card-profiles", "retrieve",
			"--digital-card-profile-id", "digital_card_profile_s3puplu90f04xhcwkiga",
		)
	})
}

func TestDigitalCardProfilesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digital-card-profiles", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(digitalCardProfilesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digital-card-profiles", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[pending]",
		)
	})
}

func TestDigitalCardProfilesArchive(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digital-card-profiles", "archive",
			"--digital-card-profile-id", "digital_card_profile_s3puplu90f04xhcwkiga",
		)
	})
}

func TestDigitalCardProfilesClone(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digital-card-profiles", "clone",
			"--digital-card-profile-id", "digital_card_profile_s3puplu90f04xhcwkiga",
			"--app-icon-file-id", "app_icon_file_id",
			"--background-image-file-id", "file_1ai913suu1zfn1pdetru",
			"--card-description", "x",
			"--contact-email", "dev@stainless.com",
			"--contact-phone", "x",
			"--contact-website", "contact_website",
			"--description", "x",
			"--issuer-name", "x",
			"--text-color", "{blue: 0, green: 0, red: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(digitalCardProfilesClone)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digital-card-profiles", "clone",
			"--digital-card-profile-id", "digital_card_profile_s3puplu90f04xhcwkiga",
			"--app-icon-file-id", "app_icon_file_id",
			"--background-image-file-id", "file_1ai913suu1zfn1pdetru",
			"--card-description", "x",
			"--contact-email", "dev@stainless.com",
			"--contact-phone", "x",
			"--contact-website", "contact_website",
			"--description", "x",
			"--issuer-name", "x",
			"--text-color.blue", "0",
			"--text-color.green", "0",
			"--text-color.red", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"app_icon_file_id: app_icon_file_id\n" +
			"background_image_file_id: file_1ai913suu1zfn1pdetru\n" +
			"card_description: x\n" +
			"contact_email: dev@stainless.com\n" +
			"contact_phone: x\n" +
			"contact_website: contact_website\n" +
			"description: x\n" +
			"issuer_name: x\n" +
			"text_color:\n" +
			"  blue: 0\n" +
			"  green: 0\n" +
			"  red: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"digital-card-profiles", "clone",
			"--digital-card-profile-id", "digital_card_profile_s3puplu90f04xhcwkiga",
		)
	})
}
