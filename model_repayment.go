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

type Repayment struct {
	// Loan record ID
	Id string `json:"id,omitempty"`
	// Repayment time
	CreateTime string `json:"create_time,omitempty"`
	// Repaid principal
	Principal string `json:"principal,omitempty"`
	// Repaid interest
	Interest string `json:"interest,omitempty"`
}
