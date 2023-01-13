/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Futures order details
type FuturesPriceTriggeredOrder struct {
	Initial     FuturesInitialOrder `json:"initial"`
	Trigger     FuturesPriceTrigger `json:"trigger"`
	StopTrigger FuturesStopTrigger  `json:"stop_trggier"`
	// Auto order ID
	Id int64 `json:"id,omitempty"`
	// User ID
	User int32 `json:"user,omitempty"`
	// Creation time
	CreateTime float64 `json:"create_time,omitempty"`
	// Finished time
	FinishTime float64 `json:"finish_time,omitempty"`
	// ID of the newly created order on condition triggered
	TradeId int64 `json:"trade_id,omitempty"`
	// Auto order status  - `open`: order is active - `finished`: order is finished - `inactive`: order is not active, only for close-long-order or close-short-order - `invalid`: order is invalid, only for close-long-order or close-short-order
	Status string `json:"status,omitempty"`
	// How order is finished
	FinishAs string `json:"finish_as,omitempty"`
	// Additional remarks on how the order was finished
	Reason string `json:"reason,omitempty"`
	// Take-profit/stop-loss types, which include:  - `close-long-order`: order take-profit/stop-loss, close long position - `close-short-order`: order take-profit/stop-loss, close short position - `close-long-position`: position take-profit/stop-loss, close long position - `close-short-position`: position take-profit/stop-loss, close short position - `plan-close-long-position`: position planned take-profit/stop-loss, close long position - `plan-close-short-position`: position planned take-profit/stop-loss, close short position  The order take-profit/stop-loss can not be passed by request. These two types are read only.
	OrderType string `json:"order_type,omitempty"`
	// Corresponding order ID of order take-profit/stop-loss.
	MeOrderId   int64 `json:"me_order_id,omitempty"`
	IsStopOrder bool  `json:"is_stop_order,omitempty"`
}
