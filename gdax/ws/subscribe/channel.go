package gdax

import bm "github.com/prodatalab/msg/basemsg"

//go:generate zebrapack -fast-strings -no-structnames-onwire -write-schema channel.schema

// Channel blah
type Channel struct {
	bm.BaseMsg `json:"basemsg" zid:"0"`
	Name       string   `json:"name" zid:"1"`
	ProductIDs []string `json:"product_ids" zid:"2"`
}
