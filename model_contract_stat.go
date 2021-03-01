/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type ContractStat struct {
	// Stat timestamp
	Time int64 `json:"time,omitempty"`
	// Long/short account number ratio
	LsrTaker float32 `json:"lsr_taker,omitempty"`
	// Long/short taker size ratio
	LsrAccount float32 `json:"lsr_account,omitempty"`
	// Long liquidation size
	LongLiqSize int64 `json:"long_liq_size,omitempty"`
	// Long liquidation amount(base currency)
	LongLiqAmount float64 `json:"long_liq_amount,omitempty"`
	// Long liquidation volume(quote currency)
	LongLiqUsd float64 `json:"long_liq_usd,omitempty"`
	// Short liquidation size
	ShortLiqSize int64 `json:"short_liq_size,omitempty"`
	// Short liquidation amount(base currency)
	ShortLiqAmount float64 `json:"short_liq_amount,omitempty"`
	// Short liquidation volume(quote currency)
	ShortLiqUsd float64 `json:"short_liq_usd,omitempty"`
	// Open interest size
	OpenInterest int64 `json:"open_interest,omitempty"`
	// Open interest volume(quote currency)
	OpenInterestUsd float64 `json:"open_interest_usd,omitempty"`
	// Top trader long/short account ratio
	TopLsrAccount float64 `json:"top_lsr_account,omitempty"`
	// Top trader long/short position ratio
	TopLsrSize float64 `json:"top_lsr_size,omitempty"`
}
