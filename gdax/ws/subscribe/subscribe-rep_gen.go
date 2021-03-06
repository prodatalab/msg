// Code generated by ZEBRAPACK (github.com/glycerine/zebrapack). DO NOT EDIT.

package gdax

import (
	"github.com/glycerine/zebrapack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *SubscribeRep) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields1zgensym_9c3aba0c5b31d397_2 = 3

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields1zgensym_9c3aba0c5b31d397_2 uint32
	totalEncodedFields1zgensym_9c3aba0c5b31d397_2, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft1zgensym_9c3aba0c5b31d397_2 := totalEncodedFields1zgensym_9c3aba0c5b31d397_2
	missingFieldsLeft1zgensym_9c3aba0c5b31d397_2 := maxFields1zgensym_9c3aba0c5b31d397_2 - totalEncodedFields1zgensym_9c3aba0c5b31d397_2

	var nextMiss1zgensym_9c3aba0c5b31d397_2 int = -1
	var found1zgensym_9c3aba0c5b31d397_2 [maxFields1zgensym_9c3aba0c5b31d397_2]bool
	var curField1zgensym_9c3aba0c5b31d397_2 int

doneWithStruct1zgensym_9c3aba0c5b31d397_2:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zgensym_9c3aba0c5b31d397_2 > 0 || missingFieldsLeft1zgensym_9c3aba0c5b31d397_2 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zgensym_9c3aba0c5b31d397_2, missingFieldsLeft1zgensym_9c3aba0c5b31d397_2, msgp.ShowFound(found1zgensym_9c3aba0c5b31d397_2[:]), decodeMsgFieldOrder1zgensym_9c3aba0c5b31d397_2)
		if encodedFieldsLeft1zgensym_9c3aba0c5b31d397_2 > 0 {
			encodedFieldsLeft1zgensym_9c3aba0c5b31d397_2--
			curField1zgensym_9c3aba0c5b31d397_2, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zgensym_9c3aba0c5b31d397_2 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss1zgensym_9c3aba0c5b31d397_2 = 0
			}
			for nextMiss1zgensym_9c3aba0c5b31d397_2 < maxFields1zgensym_9c3aba0c5b31d397_2 && (found1zgensym_9c3aba0c5b31d397_2[nextMiss1zgensym_9c3aba0c5b31d397_2] || decodeMsgFieldSkip1zgensym_9c3aba0c5b31d397_2[nextMiss1zgensym_9c3aba0c5b31d397_2]) {
				nextMiss1zgensym_9c3aba0c5b31d397_2++
			}
			if nextMiss1zgensym_9c3aba0c5b31d397_2 == maxFields1zgensym_9c3aba0c5b31d397_2 {
				// filled all the empty fields!
				break doneWithStruct1zgensym_9c3aba0c5b31d397_2
			}
			missingFieldsLeft1zgensym_9c3aba0c5b31d397_2--
			curField1zgensym_9c3aba0c5b31d397_2 = nextMiss1zgensym_9c3aba0c5b31d397_2
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zgensym_9c3aba0c5b31d397_2)
		switch curField1zgensym_9c3aba0c5b31d397_2 {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "BaseMsg"
			found1zgensym_9c3aba0c5b31d397_2[0] = true
			err = z.BaseMsg.DecodeMsg(dc)
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Type"
			found1zgensym_9c3aba0c5b31d397_2[1] = true
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Channels"
			found1zgensym_9c3aba0c5b31d397_2[2] = true
			var zgensym_9c3aba0c5b31d397_3 uint32
			zgensym_9c3aba0c5b31d397_3, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Channels) >= int(zgensym_9c3aba0c5b31d397_3) {
				z.Channels = (z.Channels)[:zgensym_9c3aba0c5b31d397_3]
			} else {
				z.Channels = make([]*Channel, zgensym_9c3aba0c5b31d397_3)
			}
			for zgensym_9c3aba0c5b31d397_0 := range z.Channels {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Channels[zgensym_9c3aba0c5b31d397_0] != nil {
						dc.PushAlwaysNil()
						err = z.Channels[zgensym_9c3aba0c5b31d397_0].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Channels[zgensym_9c3aba0c5b31d397_0] == nil {
						z.Channels[zgensym_9c3aba0c5b31d397_0] = new(Channel)
					}
					err = z.Channels[zgensym_9c3aba0c5b31d397_0].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss1zgensym_9c3aba0c5b31d397_2 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of SubscribeRep
var decodeMsgFieldOrder1zgensym_9c3aba0c5b31d397_2 = []string{"BaseMsg", "Type", "Channels"}

var decodeMsgFieldSkip1zgensym_9c3aba0c5b31d397_2 = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *SubscribeRep) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = false
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.Type) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Channels) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *SubscribeRep) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zgensym_9c3aba0c5b31d397_4 [3]bool
	fieldsInUse_zgensym_9c3aba0c5b31d397_5 := z.fieldsNotEmpty(empty_zgensym_9c3aba0c5b31d397_4[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_9c3aba0c5b31d397_5)
	if err != nil {
		return err
	}

	if !empty_zgensym_9c3aba0c5b31d397_4[0] {
		// zid 0 for "BaseMsg"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = z.BaseMsg.EncodeMsg(en)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_9c3aba0c5b31d397_4[1] {
		// zid 1 for "Type"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Type)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_9c3aba0c5b31d397_4[2] {
		// zid 2 for "Channels"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Channels)))
		if err != nil {
			return
		}
		for zgensym_9c3aba0c5b31d397_0 := range z.Channels {
			if z.Channels[zgensym_9c3aba0c5b31d397_0] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Channels[zgensym_9c3aba0c5b31d397_0].EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SubscribeRep) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "BaseMsg"
		o = append(o, 0x0)
		o, err = z.BaseMsg.MarshalMsg(o) // not is.iface, gen/marshal.go:261
		if err != nil {
			return
		}
	}

	if !empty[1] {
		// zid 1 for "Type"
		o = append(o, 0x1)
		o = msgp.AppendString(o, z.Type)
	}

	if !empty[2] {
		// zid 2 for "Channels"
		o = append(o, 0x2)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Channels)))
		for zgensym_9c3aba0c5b31d397_0 := range z.Channels {
			if z.Channels[zgensym_9c3aba0c5b31d397_0] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Channels[zgensym_9c3aba0c5b31d397_0].MarshalMsg(o) // not is.iface, gen/marshal.go:261
				if err != nil {
					return
				}
			}
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SubscribeRep) UnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.UnmarshalMsgWithCfg(bts, cfg)
}
func (z *SubscribeRep) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields6zgensym_9c3aba0c5b31d397_7 = 3

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields6zgensym_9c3aba0c5b31d397_7 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zgensym_9c3aba0c5b31d397_7, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft6zgensym_9c3aba0c5b31d397_7 := totalEncodedFields6zgensym_9c3aba0c5b31d397_7
	missingFieldsLeft6zgensym_9c3aba0c5b31d397_7 := maxFields6zgensym_9c3aba0c5b31d397_7 - totalEncodedFields6zgensym_9c3aba0c5b31d397_7

	var nextMiss6zgensym_9c3aba0c5b31d397_7 int = -1
	var found6zgensym_9c3aba0c5b31d397_7 [maxFields6zgensym_9c3aba0c5b31d397_7]bool
	var curField6zgensym_9c3aba0c5b31d397_7 int

doneWithStruct6zgensym_9c3aba0c5b31d397_7:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zgensym_9c3aba0c5b31d397_7 > 0 || missingFieldsLeft6zgensym_9c3aba0c5b31d397_7 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zgensym_9c3aba0c5b31d397_7, missingFieldsLeft6zgensym_9c3aba0c5b31d397_7, msgp.ShowFound(found6zgensym_9c3aba0c5b31d397_7[:]), unmarshalMsgFieldOrder6zgensym_9c3aba0c5b31d397_7)
		if encodedFieldsLeft6zgensym_9c3aba0c5b31d397_7 > 0 {
			encodedFieldsLeft6zgensym_9c3aba0c5b31d397_7--
			curField6zgensym_9c3aba0c5b31d397_7, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss6zgensym_9c3aba0c5b31d397_7 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zgensym_9c3aba0c5b31d397_7 = 0
			}
			for nextMiss6zgensym_9c3aba0c5b31d397_7 < maxFields6zgensym_9c3aba0c5b31d397_7 && (found6zgensym_9c3aba0c5b31d397_7[nextMiss6zgensym_9c3aba0c5b31d397_7] || unmarshalMsgFieldSkip6zgensym_9c3aba0c5b31d397_7[nextMiss6zgensym_9c3aba0c5b31d397_7]) {
				nextMiss6zgensym_9c3aba0c5b31d397_7++
			}
			if nextMiss6zgensym_9c3aba0c5b31d397_7 == maxFields6zgensym_9c3aba0c5b31d397_7 {
				// filled all the empty fields!
				break doneWithStruct6zgensym_9c3aba0c5b31d397_7
			}
			missingFieldsLeft6zgensym_9c3aba0c5b31d397_7--
			curField6zgensym_9c3aba0c5b31d397_7 = nextMiss6zgensym_9c3aba0c5b31d397_7
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zgensym_9c3aba0c5b31d397_7)
		switch curField6zgensym_9c3aba0c5b31d397_7 {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "BaseMsg"
			found6zgensym_9c3aba0c5b31d397_7[0] = true
			bts, err = z.BaseMsg.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Type"
			found6zgensym_9c3aba0c5b31d397_7[1] = true
			z.Type, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Channels"
			found6zgensym_9c3aba0c5b31d397_7[2] = true
			if nbs.AlwaysNil {
				(z.Channels) = (z.Channels)[:0]
			} else {

				var zgensym_9c3aba0c5b31d397_8 uint32
				zgensym_9c3aba0c5b31d397_8, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Channels) >= int(zgensym_9c3aba0c5b31d397_8) {
					z.Channels = (z.Channels)[:zgensym_9c3aba0c5b31d397_8]
				} else {
					z.Channels = make([]*Channel, zgensym_9c3aba0c5b31d397_8)
				}
				for zgensym_9c3aba0c5b31d397_0 := range z.Channels {
					if nbs.AlwaysNil {
						if z.Channels[zgensym_9c3aba0c5b31d397_0] != nil {
							z.Channels[zgensym_9c3aba0c5b31d397_0].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Channels[zgensym_9c3aba0c5b31d397_0] {
								z.Channels[zgensym_9c3aba0c5b31d397_0].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Channels[zgensym_9c3aba0c5b31d397_0] == nil {
								z.Channels[zgensym_9c3aba0c5b31d397_0] = new(Channel)
							}
							bts, err = z.Channels[zgensym_9c3aba0c5b31d397_0].UnmarshalMsg(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss6zgensym_9c3aba0c5b31d397_7 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of SubscribeRep
var unmarshalMsgFieldOrder6zgensym_9c3aba0c5b31d397_7 = []string{"BaseMsg", "Type", "Channels"}

var unmarshalMsgFieldSkip6zgensym_9c3aba0c5b31d397_7 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SubscribeRep) Msgsize() (s int) {
	s = 1 + 8 + z.BaseMsg.Msgsize() + 5 + msgp.StringPrefixSize + len(z.Type) + 9 + msgp.ArrayHeaderSize
	for zgensym_9c3aba0c5b31d397_0 := range z.Channels {
		if z.Channels[zgensym_9c3aba0c5b31d397_0] == nil {
			s += msgp.NilSize
		} else {
			s += z.Channels[zgensym_9c3aba0c5b31d397_0].Msgsize()
		}
	}
	return
}

// FileSubscribe_rep holds ZebraPack schema from file 'subscribe-rep.go'
type FileSubscribe_rep struct{}

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 585 bytes
func (FileSubscribe_rep) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x84, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0xb0, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
		0x62, 0x65, 0x2d, 0x72, 0x65, 0x70, 0x2e, 0x67, 0x6f, 0xad,
		0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b,
		0x61, 0x67, 0x65, 0xa4, 0x67, 0x64, 0x61, 0x78, 0xa7, 0x53,
		0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x81, 0xac, 0x53, 0x75,
		0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70,
		0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61,
		0x6d, 0x65, 0xac, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
		0x62, 0x65, 0x52, 0x65, 0x70, 0xa6, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x73, 0x93, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa7, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x73, 0x67, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa7, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x73, 0x67,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xaa, 0x62, 0x6d, 0x2e, 0x42, 0x61, 0x73,
		0x65, 0x4d, 0x73, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69,
		0x74, 0x69, 0x76, 0x65, 0x16, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72,
		0xaa, 0x62, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x73,
		0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4,
		0x54, 0x79, 0x70, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x54, 0x79,
		0x70, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x86, 0xa3, 0x5a, 0x69, 0x64,
		0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa8, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
		0x6c, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x43, 0x68, 0x61, 0x6e,
		0x6e, 0x65, 0x6c, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xaa, 0x5b, 0x5d,
		0x2a, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x1a, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x1a, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x53,
		0x6c, 0x69, 0x63, 0x65, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69,
		0x6e, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53,
		0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x43,
		0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0xa7, 0x49, 0x6d, 0x70,
		0x6f, 0x72, 0x74, 0x73, 0x91, 0xd9, 0x26, 0x62, 0x6d, 0x20,
		0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
		0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x6c,
		0x61, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x62, 0x61, 0x73,
		0x65, 0x6d, 0x73, 0x67, 0x22,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 733 bytes
func (FileSubscribe_rep) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"subscribe-rep.go","SourcePackage":"gdax","Structs":{"SubscribeRep":{"StructName":"SubscribeRep","Fields":[{"Zid":0,"FieldGoName":"BaseMsg","FieldTagName":"BaseMsg","FieldTypeStr":"bm.BaseMsg","FieldCategory":23,"FieldPrimitive":22,"FieldFullType":{"Kind":22,"Str":"bm.BaseMsg"}},{"Zid":1,"FieldGoName":"Type","FieldTagName":"Type","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":2,"FieldGoName":"Channels","FieldTagName":"Channels","FieldTypeStr":"[]*Channel","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Channel"}}}}]}},"Imports":["bm \"github.com/prodatalab/msg/basemsg\""]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 1836 bytes
func (FileSubscribe_rep) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "subscribe-rep.go",
    "SourcePackage": "gdax",
    "Structs": {
        "SubscribeRep": {
            "StructName": "SubscribeRep",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "BaseMsg",
                    "FieldTagName": "BaseMsg",
                    "FieldTypeStr": "bm.BaseMsg",
                    "FieldCategory": 23,
                    "FieldPrimitive": 22,
                    "FieldFullType": {
                        "Kind": 22,
                        "Str": "bm.BaseMsg"
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "Type",
                    "FieldTagName": "Type",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "Channels",
                    "FieldTagName": "Channels",
                    "FieldTypeStr": "[]*Channel",
                    "FieldCategory": 26,
                    "FieldFullType": {
                        "Kind": 26,
                        "Str": "Slice",
                        "Domain": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 22,
                                "Str": "Channel"
                            }
                        }
                    }
                }
            ]
        }
    },
    "Imports": [
        "bm \"github.com/prodatalab/msg/basemsg\""
    ]
}`)
}
