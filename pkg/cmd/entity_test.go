// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestEntitiesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "create",
			"--structure", "corporation",
			"--corporation", "{address: {city: New York, country: x, line1: 33 Liberty Street, line2: x, state: NY, zip: '10045'}, beneficial_owners: [{individual: {address: {city: New York, country: x, line1: 33 Liberty Street, line2: x, state: NY, zip: '10045'}, date_of_birth: '1970-01-31', identification: {method: social_security_number, number: '078051120', drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: Ian Crease, confirmed_no_us_tax_id: true}, prongs: [control], company_title: CEO}], legal_identifier: {value: '602214076', category: us_employer_identification_number}, name: National Phonograph Company, beneficial_ownership_exemption_reason: regulated_financial_institution, email: dev@stainless.com, incorporation_state: NY, industry_code: x, website: https://example.com}",
			"--description", "x",
			"--government-authority", "{address: {city: x, line1: x, state: xx, zip: x, line2: x}, authorized_persons: [{name: x}], category: municipality, name: x, tax_identifier: x, website: website}",
			"--joint", "{individuals: [{address: {city: x, country: x, line1: x, line2: x, state: x, zip: x}, date_of_birth: '2019-12-27', identification: {method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: x, confirmed_no_us_tax_id: true}]}",
			"--natural-person", "{address: {city: x, country: x, line1: x, line2: x, state: x, zip: x}, date_of_birth: '2019-12-27', identification: {method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: x, confirmed_no_us_tax_id: true}",
			"--risk-rating", "{rated_at: '2019-12-27T18:11:19.117Z', rating: low}",
			"--supplemental-document", "{file_id: file_makxrc67oh9l6sg7w9yc}",
			"--terms-agreement", "{agreed_at: '2019-12-27T18:11:19.117Z', ip_address: x, terms_url: x}",
			"--third-party-verification", "{reference: x, vendor: alloy}",
			"--trust", "{address: {city: x, line1: x, state: xx, zip: x, line2: x}, category: revocable, name: x, trustees: [{structure: individual, individual: {address: {city: x, country: x, line1: x, line2: x, state: x, zip: x}, date_of_birth: '2019-12-27', identification: {method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: x, confirmed_no_us_tax_id: true}}], formation_document_file_id: formation_document_file_id, formation_state: x, grantor: {address: {city: x, country: x, line1: x, line2: x, state: x, zip: x}, date_of_birth: '2019-12-27', identification: {method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: x, confirmed_no_us_tax_id: true}, tax_identifier: x}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(entitiesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "create",
			"--structure", "corporation",
			"--corporation.address", "{city: New York, country: x, line1: 33 Liberty Street, line2: x, state: NY, zip: '10045'}",
			"--corporation.beneficial-owners", "[{individual: {address: {city: New York, country: x, line1: 33 Liberty Street, line2: x, state: NY, zip: '10045'}, date_of_birth: '1970-01-31', identification: {method: social_security_number, number: '078051120', drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: Ian Crease, confirmed_no_us_tax_id: true}, prongs: [control], company_title: CEO}]",
			"--corporation.legal-identifier", "{value: '602214076', category: us_employer_identification_number}",
			"--corporation.name", "National Phonograph Company",
			"--corporation.beneficial-ownership-exemption-reason", "regulated_financial_institution",
			"--corporation.email", "dev@stainless.com",
			"--corporation.incorporation-state", "NY",
			"--corporation.industry-code", "x",
			"--corporation.website", "https://example.com",
			"--description", "x",
			"--government-authority.address", "{city: x, line1: x, state: xx, zip: x, line2: x}",
			"--government-authority.authorized-persons", "[{name: x}]",
			"--government-authority.category", "municipality",
			"--government-authority.name", "x",
			"--government-authority.tax-identifier", "x",
			"--government-authority.website", "website",
			"--joint.individuals", "[{address: {city: x, country: x, line1: x, line2: x, state: x, zip: x}, date_of_birth: '2019-12-27', identification: {method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: x, confirmed_no_us_tax_id: true}]",
			"--natural-person.address", "{city: x, country: x, line1: x, line2: x, state: x, zip: x}",
			"--natural-person.date-of-birth", "2019-12-27",
			"--natural-person.identification", "{method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}",
			"--natural-person.name", "x",
			"--natural-person.confirmed-no-us-tax-id=true",
			"--risk-rating.rated-at", "2019-12-27T18:11:19.117Z",
			"--risk-rating.rating", "low",
			"--supplemental-document.file-id", "file_makxrc67oh9l6sg7w9yc",
			"--terms-agreement.agreed-at", "2019-12-27T18:11:19.117Z",
			"--terms-agreement.ip-address", "x",
			"--terms-agreement.terms-url", "x",
			"--third-party-verification.reference", "x",
			"--third-party-verification.vendor", "alloy",
			"--trust.address", "{city: x, line1: x, state: xx, zip: x, line2: x}",
			"--trust.category", "revocable",
			"--trust.name", "x",
			"--trust.trustees", "[{structure: individual, individual: {address: {city: x, country: x, line1: x, line2: x, state: x, zip: x}, date_of_birth: '2019-12-27', identification: {method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: x, confirmed_no_us_tax_id: true}}]",
			"--trust.formation-document-file-id", "formation_document_file_id",
			"--trust.formation-state", "x",
			"--trust.grantor", "{address: {city: x, country: x, line1: x, line2: x, state: x, zip: x}, date_of_birth: '2019-12-27', identification: {method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: x, confirmed_no_us_tax_id: true}",
			"--trust.tax-identifier", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"structure: corporation\n" +
			"corporation:\n" +
			"  address:\n" +
			"    city: New York\n" +
			"    country: x\n" +
			"    line1: 33 Liberty Street\n" +
			"    line2: x\n" +
			"    state: NY\n" +
			"    zip: '10045'\n" +
			"  beneficial_owners:\n" +
			"    - individual:\n" +
			"        address:\n" +
			"          city: New York\n" +
			"          country: x\n" +
			"          line1: 33 Liberty Street\n" +
			"          line2: x\n" +
			"          state: NY\n" +
			"          zip: '10045'\n" +
			"        date_of_birth: '1970-01-31'\n" +
			"        identification:\n" +
			"          method: social_security_number\n" +
			"          number: '078051120'\n" +
			"          drivers_license:\n" +
			"            expiration_date: '2019-12-27'\n" +
			"            file_id: file_id\n" +
			"            state: xx\n" +
			"            back_file_id: back_file_id\n" +
			"          other:\n" +
			"            country: x\n" +
			"            description: x\n" +
			"            file_id: file_id\n" +
			"            back_file_id: back_file_id\n" +
			"            expiration_date: '2019-12-27'\n" +
			"          passport:\n" +
			"            country: x\n" +
			"            expiration_date: '2019-12-27'\n" +
			"            file_id: file_id\n" +
			"        name: Ian Crease\n" +
			"        confirmed_no_us_tax_id: true\n" +
			"      prongs:\n" +
			"        - control\n" +
			"      company_title: CEO\n" +
			"  legal_identifier:\n" +
			"    value: '602214076'\n" +
			"    category: us_employer_identification_number\n" +
			"  name: National Phonograph Company\n" +
			"  beneficial_ownership_exemption_reason: regulated_financial_institution\n" +
			"  email: dev@stainless.com\n" +
			"  incorporation_state: NY\n" +
			"  industry_code: x\n" +
			"  website: https://example.com\n" +
			"description: x\n" +
			"government_authority:\n" +
			"  address:\n" +
			"    city: x\n" +
			"    line1: x\n" +
			"    state: xx\n" +
			"    zip: x\n" +
			"    line2: x\n" +
			"  authorized_persons:\n" +
			"    - name: x\n" +
			"  category: municipality\n" +
			"  name: x\n" +
			"  tax_identifier: x\n" +
			"  website: website\n" +
			"joint:\n" +
			"  individuals:\n" +
			"    - address:\n" +
			"        city: x\n" +
			"        country: x\n" +
			"        line1: x\n" +
			"        line2: x\n" +
			"        state: x\n" +
			"        zip: x\n" +
			"      date_of_birth: '2019-12-27'\n" +
			"      identification:\n" +
			"        method: social_security_number\n" +
			"        number: xxxx\n" +
			"        drivers_license:\n" +
			"          expiration_date: '2019-12-27'\n" +
			"          file_id: file_id\n" +
			"          state: xx\n" +
			"          back_file_id: back_file_id\n" +
			"        other:\n" +
			"          country: x\n" +
			"          description: x\n" +
			"          file_id: file_id\n" +
			"          back_file_id: back_file_id\n" +
			"          expiration_date: '2019-12-27'\n" +
			"        passport:\n" +
			"          country: x\n" +
			"          expiration_date: '2019-12-27'\n" +
			"          file_id: file_id\n" +
			"      name: x\n" +
			"      confirmed_no_us_tax_id: true\n" +
			"natural_person:\n" +
			"  address:\n" +
			"    city: x\n" +
			"    country: x\n" +
			"    line1: x\n" +
			"    line2: x\n" +
			"    state: x\n" +
			"    zip: x\n" +
			"  date_of_birth: '2019-12-27'\n" +
			"  identification:\n" +
			"    method: social_security_number\n" +
			"    number: xxxx\n" +
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
			"  name: x\n" +
			"  confirmed_no_us_tax_id: true\n" +
			"risk_rating:\n" +
			"  rated_at: '2019-12-27T18:11:19.117Z'\n" +
			"  rating: low\n" +
			"supplemental_documents:\n" +
			"  - file_id: file_makxrc67oh9l6sg7w9yc\n" +
			"terms_agreements:\n" +
			"  - agreed_at: '2019-12-27T18:11:19.117Z'\n" +
			"    ip_address: x\n" +
			"    terms_url: x\n" +
			"third_party_verification:\n" +
			"  reference: x\n" +
			"  vendor: alloy\n" +
			"trust:\n" +
			"  address:\n" +
			"    city: x\n" +
			"    line1: x\n" +
			"    state: xx\n" +
			"    zip: x\n" +
			"    line2: x\n" +
			"  category: revocable\n" +
			"  name: x\n" +
			"  trustees:\n" +
			"    - structure: individual\n" +
			"      individual:\n" +
			"        address:\n" +
			"          city: x\n" +
			"          country: x\n" +
			"          line1: x\n" +
			"          line2: x\n" +
			"          state: x\n" +
			"          zip: x\n" +
			"        date_of_birth: '2019-12-27'\n" +
			"        identification:\n" +
			"          method: social_security_number\n" +
			"          number: xxxx\n" +
			"          drivers_license:\n" +
			"            expiration_date: '2019-12-27'\n" +
			"            file_id: file_id\n" +
			"            state: xx\n" +
			"            back_file_id: back_file_id\n" +
			"          other:\n" +
			"            country: x\n" +
			"            description: x\n" +
			"            file_id: file_id\n" +
			"            back_file_id: back_file_id\n" +
			"            expiration_date: '2019-12-27'\n" +
			"          passport:\n" +
			"            country: x\n" +
			"            expiration_date: '2019-12-27'\n" +
			"            file_id: file_id\n" +
			"        name: x\n" +
			"        confirmed_no_us_tax_id: true\n" +
			"  formation_document_file_id: formation_document_file_id\n" +
			"  formation_state: x\n" +
			"  grantor:\n" +
			"    address:\n" +
			"      city: x\n" +
			"      country: x\n" +
			"      line1: x\n" +
			"      line2: x\n" +
			"      state: x\n" +
			"      zip: x\n" +
			"    date_of_birth: '2019-12-27'\n" +
			"    identification:\n" +
			"      method: social_security_number\n" +
			"      number: xxxx\n" +
			"      drivers_license:\n" +
			"        expiration_date: '2019-12-27'\n" +
			"        file_id: file_id\n" +
			"        state: xx\n" +
			"        back_file_id: back_file_id\n" +
			"      other:\n" +
			"        country: x\n" +
			"        description: x\n" +
			"        file_id: file_id\n" +
			"        back_file_id: back_file_id\n" +
			"        expiration_date: '2019-12-27'\n" +
			"      passport:\n" +
			"        country: x\n" +
			"        expiration_date: '2019-12-27'\n" +
			"        file_id: file_id\n" +
			"    name: x\n" +
			"    confirmed_no_us_tax_id: true\n" +
			"  tax_identifier: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entities", "create",
		)
	})
}

func TestEntitiesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "retrieve",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
		)
	})
}

func TestEntitiesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "update",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--corporation", "{address: {city: New York, country: US, line1: 33 Liberty Street, line2: Unit 2, state: NY, zip: '10045'}, email: dev@stainless.com, incorporation_state: x, industry_code: x, legal_identifier: {value: x, category: us_employer_identification_number}, name: x, website: website}",
			"--details-confirmed-at", "'2019-12-27T18:11:19.117Z'",
			"--government-authority", "{address: {city: x, line1: x, state: xx, zip: x, line2: x}, name: x}",
			"--natural-person", "{address: {city: x, country: x, line1: x, line2: x, state: x, zip: x}, confirmed_no_us_tax_id: true, identification: {method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}, name: x}",
			"--risk-rating", "{rated_at: '2020-01-31T23:59:59Z', rating: low}",
			"--terms-agreement", "{agreed_at: '2019-12-27T18:11:19.117Z', ip_address: x, terms_url: x}",
			"--third-party-verification", "{reference: x, vendor: alloy}",
			"--trust", "{address: {city: x, line1: x, state: xx, zip: x, line2: x}, name: x}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(entitiesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "update",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--corporation.address", "{city: New York, country: US, line1: 33 Liberty Street, line2: Unit 2, state: NY, zip: '10045'}",
			"--corporation.email", "dev@stainless.com",
			"--corporation.incorporation-state", "x",
			"--corporation.industry-code", "x",
			"--corporation.legal-identifier", "{value: x, category: us_employer_identification_number}",
			"--corporation.name", "x",
			"--corporation.website", "website",
			"--details-confirmed-at", "'2019-12-27T18:11:19.117Z'",
			"--government-authority.address", "{city: x, line1: x, state: xx, zip: x, line2: x}",
			"--government-authority.name", "x",
			"--natural-person.address", "{city: x, country: x, line1: x, line2: x, state: x, zip: x}",
			"--natural-person.confirmed-no-us-tax-id=true",
			"--natural-person.identification", "{method: social_security_number, number: xxxx, drivers_license: {expiration_date: '2019-12-27', file_id: file_id, state: xx, back_file_id: back_file_id}, other: {country: x, description: x, file_id: file_id, back_file_id: back_file_id, expiration_date: '2019-12-27'}, passport: {country: x, expiration_date: '2019-12-27', file_id: file_id}}",
			"--natural-person.name", "x",
			"--risk-rating.rated-at", "2020-01-31T23:59:59Z",
			"--risk-rating.rating", "low",
			"--terms-agreement.agreed-at", "2019-12-27T18:11:19.117Z",
			"--terms-agreement.ip-address", "x",
			"--terms-agreement.terms-url", "x",
			"--third-party-verification.reference", "x",
			"--third-party-verification.vendor", "alloy",
			"--trust.address", "{city: x, line1: x, state: xx, zip: x, line2: x}",
			"--trust.name", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"corporation:\n" +
			"  address:\n" +
			"    city: New York\n" +
			"    country: US\n" +
			"    line1: 33 Liberty Street\n" +
			"    line2: Unit 2\n" +
			"    state: NY\n" +
			"    zip: '10045'\n" +
			"  email: dev@stainless.com\n" +
			"  incorporation_state: x\n" +
			"  industry_code: x\n" +
			"  legal_identifier:\n" +
			"    value: x\n" +
			"    category: us_employer_identification_number\n" +
			"  name: x\n" +
			"  website: website\n" +
			"details_confirmed_at: '2019-12-27T18:11:19.117Z'\n" +
			"government_authority:\n" +
			"  address:\n" +
			"    city: x\n" +
			"    line1: x\n" +
			"    state: xx\n" +
			"    zip: x\n" +
			"    line2: x\n" +
			"  name: x\n" +
			"natural_person:\n" +
			"  address:\n" +
			"    city: x\n" +
			"    country: x\n" +
			"    line1: x\n" +
			"    line2: x\n" +
			"    state: x\n" +
			"    zip: x\n" +
			"  confirmed_no_us_tax_id: true\n" +
			"  identification:\n" +
			"    method: social_security_number\n" +
			"    number: xxxx\n" +
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
			"  name: x\n" +
			"risk_rating:\n" +
			"  rated_at: '2020-01-31T23:59:59Z'\n" +
			"  rating: low\n" +
			"terms_agreements:\n" +
			"  - agreed_at: '2019-12-27T18:11:19.117Z'\n" +
			"    ip_address: x\n" +
			"    terms_url: x\n" +
			"third_party_verification:\n" +
			"  reference: x\n" +
			"  vendor: alloy\n" +
			"trust:\n" +
			"  address:\n" +
			"    city: x\n" +
			"    line1: x\n" +
			"    state: xx\n" +
			"    zip: x\n" +
			"    line2: x\n" +
			"  name: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entities", "update",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
		)
	})
}

func TestEntitiesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "list",
			"--max-items", "10",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [active]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(entitiesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "list",
			"--max-items", "10",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[active]",
		)
	})
}

func TestEntitiesArchive(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "archive",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
		)
	})
}
