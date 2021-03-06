package rpc

import (
	"gopkg.in/vmihailenco/msgpack.v2"
)

type MsgPackCodec struct {
}

func (c *MsgPackCodec) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (c *MsgPackCodec) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

func (c *MsgPackCodec) String() string {
	return "MsgPackCodec"
}
