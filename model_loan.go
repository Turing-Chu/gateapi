/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.6.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type Loan struct {
	// Loan ID
	Id string `json:"id,omitempty"`
	// Creation time
	CreateTime string `json:"create_time,omitempty"`
	// Repay time of the loan. No value will be returned for lending loan
	ExpireTime string `json:"expire_time,omitempty"`
	// Loan status  open - not fully loaned loaned - all loaned out for lending loan; loaned in for borrowing side finished - loan is finished, either being all repaid or cancelled by the lender auto_repaid - automatically repaid by the system
	Status string `json:"status,omitempty"`
	// Loan side
	Side string `json:"side"`
	// Loan currency
	Currency string `json:"currency"`
	// Loan rate
	Rate string `json:"rate"`
	// Loan amount
	Amount string `json:"amount"`
	// Loan days
	Days int32 `json:"days"`
	// Auto renew the loan on expiration
	AutoRenew bool `json:"auto_renew,omitempty"`
	// Currency pair. Required for borrowing side
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Amount not lending out
	Left string `json:"left,omitempty"`
	// Repaid amount
	Repaid string `json:"repaid,omitempty"`
	// Repaid interest
	PaidInterest string `json:"paid_interest,omitempty"`
	// Interest not repaid
	UnpaidInterest string `json:"unpaid_interest,omitempty"`
}
