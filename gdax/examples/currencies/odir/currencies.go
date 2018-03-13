package gdax

type CurrenciesCurrencies []struct {
	ID      string `json:"id"`
	MinSize string `json:"min_size"`
	Name    string `json:"name"`
}
