package bytearray

//go:generate greenpack

// Base blah
type Base struct {
	Version int8  `msg:"version" zid:"0"`
	Type    int64 `msg:"type" zid:"1"`
}

// ByteArray blah
type ByteArray struct {
	Base  `zid:"0"`
	Value []byte `msg:"value" zid:"1"`
}
