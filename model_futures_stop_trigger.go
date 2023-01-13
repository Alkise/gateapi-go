package gateapi

type FuturesStopTrigger struct {
	// Trigger condition type  - `1`: calculated price based on `strategy_type` and `price_type` >= `price` - `2`: calculated price based on `strategy_type` and `price_type` <= `price`
	Rule         int32  `json:"rule,omitempty"`
	TriggerPrice string `json:"trigger_price,omitempty"`
	OrderPrice   string `json:"order_price,omitempty"`
}
