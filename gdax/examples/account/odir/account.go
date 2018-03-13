package gdax

type AccountAccount struct {
	Available string `json:"available" zid:"0"`
	Balance   string `json:"balance" zid:"1"`
	Currency  string `json:"currency" zid:"2"`
	Hold      string `json:"hold" zid:"3"`
	ID        string `json:"id" zid:"4"`
	ProfileID string `json:"profile_id" zid:"5"`
}
