package gateapi

type SubAccountInnerTransfer struct {
	// Transfer currency name
	Currency string `json:"currency"`
	// Transfer from the account. (deprecate, use sub_account_from_type and sub_account_to_type instead)
	SubAccountType string `json:"sub_account_type,omitempty"`
	// From sub user's account. Transfer from the user id of the sub-account
	SubAccountFrom string `json:"sub_account_from"`
	// From sub user's account. `spot` - spot account, `futures` - perpetual contract account, `cross_margin` - cross margin account
	SubAccountFromType string `json:"sub_account_from_type"`
	// Target sub user's account. Transfer from the user id of the sub-account
	SubAccountTo string `json:"sub_account_to"`
	// Target sub user's account. `spot` - spot account, `futures` - perpetual contract account, `cross_margin` - cross margin account
	SubAccountToType string `json:"sub_account_to_type"`
	// Transfer amount
	Amount string `json:"amount"`
}
