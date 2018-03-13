package msg

//go:generate zebrapack

// ByteArray blah
type ByteArray struct {
	BaseMsg `json:"basemsg" zid:"0"`
	Value   []byte `json:"value" zid:"1"`
}
