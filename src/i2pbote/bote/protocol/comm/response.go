package comm

import (
	"i2pbote/bote/common"
	"i2pbote/util"
)

type ResponseStatus byte

const OK = ResponseStatus(0)
const Err = ResponseStatus(1)
const NoData = ResponseStatus(2)
const InvalidPacket = ResponseStatus(3)
const InvalidHashcash = ResponseStatus(4)
const NotEnoughHashCash = ResponseStatus(5)

func (s ResponseStatus) Byte() byte {
	return byte(s)
}

// create response packet
func ResponsePacket(version byte, status ResponseStatus, cid common.CID, data []byte) *Packet {
	if data == nil {
		data = []byte{}
	}
	l := len(data)
	buff := make([]byte, 1+32+2+l)
	buff[0] = status.Byte()
	copy(buff[1:], cid[:])
	util.PutUInt16_i(l, buff[33:])
	if l > 0 {
		copy(buff[35:], data)
	}
	return New(version, Response, buff)
}
