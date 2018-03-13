package gdax

type AccountHistory []struct {
	Amount    string              `json:"amount"`
	Balance   string              `json:"balance"`
	CreatedAt string              `json:"created_at"`
	Details   AccountHistory_sub1 `json:"details"`
	ID        string              `json:"id"`
	Type      string              `json:"type"`
}
