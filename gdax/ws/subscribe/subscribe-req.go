package gdax

import bm "github.com/prodatalab/msg/basemsg"

type chan2 Channel

//go:generate zebrapack -no-load -fast-strings -no-structnames-onwire -write-schema subscribe-req.schema

// SubscribeReq blah
type SubscribeReq struct {
	bm.BaseMsg `json:"basemsg" msg:"basemsg" zid:"0"`
	Type       string     `json:"type" msg:"type" zid:"1"`
	Channels   []*Channel `json:"channels" msg:"channels" zid:"2"`
}
