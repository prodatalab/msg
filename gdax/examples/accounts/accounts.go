package gdax

type AccountsAccounts []struct {
	Available string `json:"available"`
	Balance   string `json:"balance"`
	Currency  string `json:"currency"`
	Hold      string `json:"hold"`
	ID        string `json:"id"`
	ProfileID string `json:"profile_id"`
}
