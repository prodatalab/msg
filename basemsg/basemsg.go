package basemsg

//go:generate zebrapack -fast-strings -no-structnames-onwire -write-schema basemsg.schema

const zebraSchemaId64 int64 = 0x1e796f175a9f5 // 536110458841589

// BaseMsg blah
type BaseMsg struct {
	APIType    int64 `json:"api_type" zid:"1"`
	APIVersion int8  `json:"api_version" zid:"0"`
}
