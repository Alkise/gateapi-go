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

type Trade struct {
	// Trade ID
	Id string `json:"id,omitempty"`
	// Trading time
	CreateTime string `json:"create_time,omitempty"`
	// Order side
	Side string `json:"side,omitempty"`
	// Trade amount
	Amount string `json:"amount,omitempty"`
	// Order price
	Price string `json:"price,omitempty"`
	// Related order ID. No value in public endpoints
	OrderId string `json:"order_id,omitempty"`
}
