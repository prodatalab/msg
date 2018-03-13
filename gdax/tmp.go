package gdax

type SubscribeReq struct {
	Channels   []interface{} `json:"channels"`
	ProductIds []string      `json:"product_ids"`
	Type       string        `json:"type"`
}
