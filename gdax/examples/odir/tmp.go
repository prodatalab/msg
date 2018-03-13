package gdax

type SubscribeReq struct {
	Channels   []interface{} `json:"channels"`
	ProductIds []string      `json:"product_ids" zid:"0"`
	Type       string        `json:"type" zid:"1"`
}
