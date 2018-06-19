package gdax

import bm "github.com/prodatalab/msg/basemsg"

//go:generate zebrapack -no-load -fast-strings -no-structnames-onwire -write-schema subscribe-rep.schema

// SubscribeRep blah
type SubscribeRep struct {
	bm.BaseMsg `json:"basemsg" msg:"basemsg" zid:"0"`
	Type       string     `json:"type" msg:"type" zid:"1"`
	Channels   []*Channel `json:"channels" msg:"channels" zid:"2"`
}
