package gdax

type AccountsCoinbase []struct {
	Active                 bool                  `json:"active"`
	Balance                string                `json:"balance"`
	Currency               string                `json:"currency"`
	ID                     string                `json:"id"`
	Name                   string                `json:"name"`
	Primary                bool                  `json:"primary"`
	SepaDepositInformation AccountsCoinbase_sub1 `json:"sepa_deposit_information"`
	Type                   string                `json:"type"`
	WireDepositInformation AccountsCoinbase_sub3 `json:"wire_deposit_information"`
}
