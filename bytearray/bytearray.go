package bytearray

import bm "github.com/prodatalab/msg/basemsg"

//go:generate zebrapack -fast-strings -no-structnames-onwire -write-schema bytearray.schema
const zebraSchemaId64 int64 = 0x86f9e3b61324 // 148407825339172

// ByteArray blah
type ByteArray struct {
	bm.BaseMsg `json:"basemsg" zid:"0"`
	Value      []byte `json:"value" zid:"1"`
}
