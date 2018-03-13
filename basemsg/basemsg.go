package msg

//go:generate zebrapack

// BaseMsg blah
type BaseMsg struct {
	Version int8  `json:"version" zid:"0"`
	Type    int64 `json:"type" zid:"1"`
}
