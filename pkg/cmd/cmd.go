// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Increase/increase-cli/internal/autocomplete"
	"github.com/Increase/increase-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "increase",
		Usage:     "CLI for the increase API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("INCREASE_API_KEY"),
			},
			&requestflag.Flag[string]{
				Name:    "webhook-secret",
				Sources: cli.EnvVars("INCREASE_WEBHOOK_SECRET"),
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "Set the environment for API requests",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsCreate,
					&accountsRetrieve,
					&accountsUpdate,
					&accountsList,
					&accountsBalance,
					&accountsClose,
				},
			},
			{
				Name:     "account-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountNumbersCreate,
					&accountNumbersRetrieve,
					&accountNumbersUpdate,
					&accountNumbersList,
				},
			},
			{
				Name:     "account-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountTransfersCreate,
					&accountTransfersRetrieve,
					&accountTransfersList,
					&accountTransfersApprove,
					&accountTransfersCancel,
				},
			},
			{
				Name:     "cards",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardsCreate,
					&cardsRetrieve,
					&cardsUpdate,
					&cardsList,
					&cardsCreateDetailsIframe,
					&cardsDetails,
					&cardsUpdatePin,
				},
			},
			{
				Name:     "card-payments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardPaymentsRetrieve,
					&cardPaymentsList,
				},
			},
			{
				Name:     "card-purchase-supplements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardPurchaseSupplementsRetrieve,
					&cardPurchaseSupplementsList,
				},
			},
			{
				Name:     "card-disputes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardDisputesCreate,
					&cardDisputesRetrieve,
					&cardDisputesList,
					&cardDisputesSubmitUserSubmission,
					&cardDisputesWithdraw,
				},
			},
			{
				Name:     "physical-cards",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&physicalCardsCreate,
					&physicalCardsRetrieve,
					&physicalCardsUpdate,
					&physicalCardsList,
				},
			},
			{
				Name:     "digital-card-profiles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&digitalCardProfilesCreate,
					&digitalCardProfilesRetrieve,
					&digitalCardProfilesList,
					&digitalCardProfilesArchive,
					&digitalCardProfilesClone,
				},
			},
			{
				Name:     "physical-card-profiles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&physicalCardProfilesCreate,
					&physicalCardProfilesRetrieve,
					&physicalCardProfilesList,
					&physicalCardProfilesArchive,
					&physicalCardProfilesClone,
				},
			},
			{
				Name:     "digital-wallet-tokens",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&digitalWalletTokensRetrieve,
					&digitalWalletTokensList,
				},
			},
			{
				Name:     "transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transactionsRetrieve,
					&transactionsList,
				},
			},
			{
				Name:     "pending-transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&pendingTransactionsCreate,
					&pendingTransactionsRetrieve,
					&pendingTransactionsList,
					&pendingTransactionsRelease,
				},
			},
			{
				Name:     "declined-transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&declinedTransactionsRetrieve,
					&declinedTransactionsList,
				},
			},
			{
				Name:     "ach-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&achTransfersCreate,
					&achTransfersRetrieve,
					&achTransfersList,
					&achTransfersApprove,
					&achTransfersCancel,
				},
			},
			{
				Name:     "inbound-ach-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundACHTransfersRetrieve,
					&inboundACHTransfersList,
					&inboundACHTransfersCreateNotificationOfChange,
					&inboundACHTransfersDecline,
					&inboundACHTransfersTransferReturn,
				},
			},
			{
				Name:     "ach-prenotifications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&achPrenotificationsCreate,
					&achPrenotificationsRetrieve,
					&achPrenotificationsList,
				},
			},
			{
				Name:     "wire-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wireTransfersCreate,
					&wireTransfersRetrieve,
					&wireTransfersList,
					&wireTransfersApprove,
					&wireTransfersCancel,
				},
			},
			{
				Name:     "inbound-wire-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundWireTransfersRetrieve,
					&inboundWireTransfersList,
					&inboundWireTransfersReverse,
				},
			},
			{
				Name:     "wire-drawdown-requests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wireDrawdownRequestsCreate,
					&wireDrawdownRequestsRetrieve,
					&wireDrawdownRequestsList,
				},
			},
			{
				Name:     "inbound-wire-drawdown-requests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundWireDrawdownRequestsRetrieve,
					&inboundWireDrawdownRequestsList,
				},
			},
			{
				Name:     "check-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&checkTransfersCreate,
					&checkTransfersRetrieve,
					&checkTransfersList,
					&checkTransfersApprove,
					&checkTransfersCancel,
					&checkTransfersStopPayment,
				},
			},
			{
				Name:     "inbound-check-deposits",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundCheckDepositsRetrieve,
					&inboundCheckDepositsList,
					&inboundCheckDepositsDecline,
					&inboundCheckDepositsReturn,
				},
			},
			{
				Name:     "real-time-payments-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&realTimePaymentsTransfersCreate,
					&realTimePaymentsTransfersRetrieve,
					&realTimePaymentsTransfersList,
					&realTimePaymentsTransfersApprove,
					&realTimePaymentsTransfersCancel,
				},
			},
			{
				Name:     "inbound-real-time-payments-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundRealTimePaymentsTransfersRetrieve,
					&inboundRealTimePaymentsTransfersList,
				},
			},
			{
				Name:     "fednow-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&fednowTransfersCreate,
					&fednowTransfersRetrieve,
					&fednowTransfersList,
					&fednowTransfersApprove,
					&fednowTransfersCancel,
				},
			},
			{
				Name:     "inbound-fednow-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundFednowTransfersRetrieve,
					&inboundFednowTransfersList,
				},
			},
			{
				Name:     "swift-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&swiftTransfersCreate,
					&swiftTransfersRetrieve,
					&swiftTransfersList,
					&swiftTransfersApprove,
					&swiftTransfersCancel,
				},
			},
			{
				Name:     "check-deposits",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&checkDepositsCreate,
					&checkDepositsRetrieve,
					&checkDepositsList,
				},
			},
			{
				Name:     "lockbox-addresses",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&lockboxAddressesCreate,
					&lockboxAddressesRetrieve,
					&lockboxAddressesUpdate,
					&lockboxAddressesList,
				},
			},
			{
				Name:     "lockbox-recipients",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&lockboxRecipientsCreate,
					&lockboxRecipientsRetrieve,
					&lockboxRecipientsUpdate,
					&lockboxRecipientsList,
				},
			},
			{
				Name:     "inbound-mail-items",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundMailItemsRetrieve,
					&inboundMailItemsList,
					&inboundMailItemsAction,
				},
			},
			{
				Name:     "routing-numbers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&routingNumbersList,
				},
			},
			{
				Name:     "external-accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&externalAccountsCreate,
					&externalAccountsRetrieve,
					&externalAccountsUpdate,
					&externalAccountsList,
				},
			},
			{
				Name:     "entities",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&entitiesCreate,
					&entitiesRetrieve,
					&entitiesUpdate,
					&entitiesList,
					&entitiesArchive,
				},
			},
			{
				Name:     "beneficial-owners",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&beneficialOwnersCreate,
					&beneficialOwnersRetrieve,
					&beneficialOwnersUpdate,
					&beneficialOwnersList,
					&beneficialOwnersArchive,
				},
			},
			{
				Name:     "supplemental-documents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&supplementalDocumentsCreate,
					&supplementalDocumentsList,
				},
			},
			{
				Name:     "entity-onboarding-sessions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&entityOnboardingSessionsCreate,
					&entityOnboardingSessionsRetrieve,
					&entityOnboardingSessionsList,
					&entityOnboardingSessionsExpire,
				},
			},
			{
				Name:     "programs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&programsRetrieve,
					&programsList,
				},
			},
			{
				Name:     "account-statements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountStatementsRetrieve,
					&accountStatementsList,
				},
			},
			{
				Name:     "files",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&filesCreate,
					&filesRetrieve,
					&filesList,
				},
			},
			{
				Name:     "file-links",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&fileLinksCreate,
				},
			},
			{
				Name:     "exports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&exportsCreate,
					&exportsRetrieve,
					&exportsList,
				},
			},
			{
				Name:     "events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&eventsRetrieve,
					&eventsList,
				},
			},
			{
				Name:     "event-subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&eventSubscriptionsCreate,
					&eventSubscriptionsRetrieve,
					&eventSubscriptionsUpdate,
					&eventSubscriptionsList,
				},
			},
			{
				Name:     "real-time-decisions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&realTimeDecisionsRetrieve,
					&realTimeDecisionsAction,
				},
			},
			{
				Name:     "groups",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&groupsRetrieve,
				},
			},
			{
				Name:     "oauth-applications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&oauthApplicationsRetrieve,
					&oauthApplicationsList,
				},
			},
			{
				Name:     "oauth-connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&oauthConnectionsRetrieve,
					&oauthConnectionsList,
				},
			},
			{
				Name:     "oauth-tokens",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&oauthTokensCreate,
				},
			},
			{
				Name:     "intrafi-account-enrollments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&intrafiAccountEnrollmentsCreate,
					&intrafiAccountEnrollmentsRetrieve,
					&intrafiAccountEnrollmentsList,
					&intrafiAccountEnrollmentsUnenroll,
				},
			},
			{
				Name:     "intrafi-balances",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&intrafiBalancesIntrafiBalance,
				},
			},
			{
				Name:     "intrafi-exclusions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&intrafiExclusionsCreate,
					&intrafiExclusionsRetrieve,
					&intrafiExclusionsList,
					&intrafiExclusionsArchive,
				},
			},
			{
				Name:     "card-tokens",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardTokensRetrieve,
					&cardTokensList,
					&cardTokensCapabilities,
				},
			},
			{
				Name:     "card-push-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardPushTransfersCreate,
					&cardPushTransfersRetrieve,
					&cardPushTransfersList,
					&cardPushTransfersApprove,
					&cardPushTransfersCancel,
				},
			},
			{
				Name:     "card-validations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardValidationsCreate,
					&cardValidationsRetrieve,
					&cardValidationsList,
				},
			},
			{
				Name:     "simulations:interest-payments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsInterestPaymentsCreate,
				},
			},
			{
				Name:     "simulations:account-revenue-payments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsAccountRevenuePaymentsCreate,
				},
			},
			{
				Name:     "simulations:card-authorizations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardAuthorizationsCreate,
				},
			},
			{
				Name:     "simulations:card-balance-inquiries",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardBalanceInquiriesCreate,
				},
			},
			{
				Name:     "simulations:card-authorization-expirations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardAuthorizationExpirationsCreate,
				},
			},
			{
				Name:     "simulations:card-settlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardSettlementsCreate,
				},
			},
			{
				Name:     "simulations:card-reversals",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardReversalsCreate,
				},
			},
			{
				Name:     "simulations:card-increments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardIncrementsCreate,
				},
			},
			{
				Name:     "simulations:card-fuel-confirmations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardFuelConfirmationsCreate,
				},
			},
			{
				Name:     "simulations:card-refunds",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardRefundsCreate,
				},
			},
			{
				Name:     "simulations:card-authentications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardAuthenticationsCreate,
					&simulationsCardAuthenticationsChallengeAttempts,
					&simulationsCardAuthenticationsChallenges,
				},
			},
			{
				Name:     "simulations:card-purchase-supplements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardPurchaseSupplementsCreate,
				},
			},
			{
				Name:     "simulations:card-disputes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardDisputesAction,
				},
			},
			{
				Name:     "simulations:physical-cards",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsPhysicalCardsCreate,
					&simulationsPhysicalCardsAdvanceShipment,
				},
			},
			{
				Name:     "simulations:digital-wallet-token-requests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsDigitalWalletTokenRequestsCreate,
				},
			},
			{
				Name:     "simulations:pending-transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsPendingTransactionsReleaseInboundFundsHold,
				},
			},
			{
				Name:     "simulations:ach-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsACHTransfersAcknowledge,
					&simulationsACHTransfersCreateNotificationOfChange,
					&simulationsACHTransfersReturn,
					&simulationsACHTransfersSettle,
					&simulationsACHTransfersSubmit,
				},
			},
			{
				Name:     "simulations:inbound-ach-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsInboundACHTransfersCreate,
				},
			},
			{
				Name:     "simulations:wire-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsWireTransfersReverse,
					&simulationsWireTransfersSubmit,
				},
			},
			{
				Name:     "simulations:inbound-wire-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsInboundWireTransfersCreate,
				},
			},
			{
				Name:     "simulations:wire-drawdown-requests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsWireDrawdownRequestsRefuse,
					&simulationsWireDrawdownRequestsSubmit,
				},
			},
			{
				Name:     "simulations:inbound-wire-drawdown-requests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsInboundWireDrawdownRequestsCreate,
				},
			},
			{
				Name:     "simulations:check-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCheckTransfersMail,
				},
			},
			{
				Name:     "simulations:inbound-check-deposits",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsInboundCheckDepositsCreate,
					&simulationsInboundCheckDepositsAdjustment,
				},
			},
			{
				Name:     "simulations:real-time-payments-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsRealTimePaymentsTransfersComplete,
				},
			},
			{
				Name:     "simulations:inbound-real-time-payments-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsInboundRealTimePaymentsTransfersCreate,
				},
			},
			{
				Name:     "simulations:inbound-fednow-transfers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsInboundFednowTransfersCreate,
				},
			},
			{
				Name:     "simulations:check-deposits",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCheckDepositsAdjustment,
					&simulationsCheckDepositsReject,
					&simulationsCheckDepositsReturn,
					&simulationsCheckDepositsSubmit,
				},
			},
			{
				Name:     "simulations:inbound-mail-items",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsInboundMailItemsCreate,
				},
			},
			{
				Name:     "simulations:entities",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsEntitiesUpdateValidation,
				},
			},
			{
				Name:     "simulations:entity-onboarding-sessions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsEntityOnboardingSessionsSubmit,
				},
			},
			{
				Name:     "simulations:programs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsProgramsCreate,
				},
			},
			{
				Name:     "simulations:account-statements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsAccountStatementsCreate,
				},
			},
			{
				Name:     "simulations:exports",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsExportsCreate,
				},
			},
			{
				Name:     "simulations:card-tokens",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simulationsCardTokensCreate,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "increase @manpages [-o increase.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "increase.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "increase.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
