// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestBeneficialOwnersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beneficial-owners", "create",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--individual", "{address: {city: New York, country: US, line1: 33 Liberty Street, line2: x, state: NY, zip: '10045'}, date_of_birth: '1970-01-31', identification: {method: social_security_number, number: '078051120', drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: Ian Crease, confirmed_no_us_tax_id: true}",
			"--prong", "control",
			"--company-title", "CEO",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(beneficialOwnersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beneficial-owners", "create",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--individual.address", "{city: New York, country: US, line1: 33 Liberty Street, line2: x, state: NY, zip: '10045'}",
			"--individual.date-of-birth", "1970-01-31",
			"--individual.identification", "{method: social_security_number, number: '078051120', drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}",
			"--individual.name", "Ian Crease",
			"--individual.confirmed-no-us-tax-id=true",
			"--prong", "control",
			"--company-title", "CEO",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entity_id: entity_n8y8tnk2p9339ti393yi\n" +
			"individual:\n" +
			"  address:\n" +
			"    city: New York\n" +
			"    country: US\n" +
			"    line1: 33 Liberty Street\n" +
			"    line2: x\n" +
			"    state: NY\n" +
			"    zip: '10045'\n" +
			"  date_of_birth: '1970-01-31'\n" +
			"  identification:\n" +
			"    method: social_security_number\n" +
			"    number: '078051120'\n" +
			"    drivers_license:\n" +
			"      expiration_date: '2019-12-27'\n" +
			"      file_id: file_id\n" +
			"      state: xx\n" +
			"      back_file_id: back_file_id\n" +
			"    other:\n" +
			"      country: x\n" +
			"      description: x\n" +
			"      file_id: file_id\n" +
			"      back_file_id: back_file_id\n" +
			"      expiration_date: '2019-12-27'\n" +
			"    passport:\n" +
			"      country: x\n" +
			"      expiration_date: '2019-12-27'\n" +
			"      file_id: file_id\n" +
			"  name: Ian Crease\n" +
			"  confirmed_no_us_tax_id: true\n" +
			"prongs:\n" +
			"  - control\n" +
			"company_title: CEO\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beneficial-owners", "create",
		)
	})
}

func TestBeneficialOwnersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beneficial-owners", "retrieve",
			"--entity-beneficial-owner-id", "entity_beneficial_owner_vozma8szzu1sxezp5zq6",
		)
	})
}

func TestBeneficialOwnersUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beneficial-owners", "update",
			"--entity-beneficial-owner-id", "entity_beneficial_owner_vozma8szzu1sxezp5zq6",
			"--address", "{city: New York, country: US, line1: 33 Liberty Street, line2: Unit 2, state: NY, zip: '10045'}",
			"--confirmed-no-us-tax-id=true",
			"--identification", "{method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}",
			"--name", "x",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(beneficialOwnersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beneficial-owners", "update",
			"--entity-beneficial-owner-id", "entity_beneficial_owner_vozma8szzu1sxezp5zq6",
			"--address.city", "New York",
			"--address.country", "US",
			"--address.line1", "33 Liberty Street",
			"--address.line2", "Unit 2",
			"--address.state", "NY",
			"--address.zip", "10045",
			"--confirmed-no-us-tax-id=true",
			"--identification.method", "social_security_number",
			"--identification.number", "xxxx",
			"--identification.drivers-license", "{expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}",
			"--identification.other", "{country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}",
			"--identification.passport", "{country: x, expiration_date: '2019-12-27', file_id: file_id}",
			"--name", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address:\n" +
			"  city: New York\n" +
			"  country: US\n" +
			"  line1: 33 Liberty Street\n" +
			"  line2: Unit 2\n" +
			"  state: NY\n" +
			"  zip: '10045'\n" +
			"confirmed_no_us_tax_id: true\n" +
			"identification:\n" +
			"  method: social_security_number\n" +
			"  number: xxxx\n" +
			"  drivers_license:\n" +
			"    expiration_date: '2019-12-27'\n" +
			"    file_id: file_id\n" +
			"    state: xx\n" +
			"    back_file_id: back_file_id\n" +
			"  other:\n" +
			"    country: x\n" +
			"    description: x\n" +
			"    file_id: file_id\n" +
			"    back_file_id: back_file_id\n" +
			"    expiration_date: '2019-12-27'\n" +
			"  passport:\n" +
			"    country: x\n" +
			"    expiration_date: '2019-12-27'\n" +
			"    file_id: file_id\n" +
			"name: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beneficial-owners", "update",
			"--entity-beneficial-owner-id", "entity_beneficial_owner_vozma8szzu1sxezp5zq6",
		)
	})
}

func TestBeneficialOwnersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beneficial-owners", "list",
			"--max-items", "10",
			"--entity-id", "entity_id",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})
}

func TestBeneficialOwnersArchive(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beneficial-owners", "archive",
			"--entity-beneficial-owner-id", "entity_beneficial_owner_vozma8szzu1sxezp5zq6",
		)
	})
}
