package geerpc

import (
	"geerpc/codec"
	"time"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber    int        // MagicNumber marks this is a geerpc request
	CodecType      codec.Type // client may choose different Codec to encode body
	ConnectTimeout time.Duration // 0 means no limit
	HandleTimeout  time.Duration
}

var DefaultOption = &Option{
	MagicNumber:    MagicNumber,
	CodecType:      codec.GobType,
	ConnectTimeout: time.Second * 10,
}
