package protocol

import (
	"bytes"
	"errors"
)

// comm packet type byte
type CommType byte

// protocol version
type ProtoVersion byte

// i2pbote packet header
var pktHeader = []byte{0x6D, 0x30, 0x52, 0xe9}

const CommFetchReq = CommType(0x46)
const CommResponse = CommType(0x4e)
const CommPeerListReq = CommType(0x41)
const CommRelayReq = CommType(0x52)

// raw communication packet
type CommPacket struct {
	Type    CommType
	Version ProtoVersion
	Body    []byte
}

var ErrTooSmall = errors.New("too small")
var ErrBadHeader = errors.New("bad packet header")

// parse a comm packet from a byteslice
func ParseCommPacket(data []byte) (pkt *CommPacket, err error) {
	if len(data) <= 6 {
		err = ErrTooSmall
		return
	}
	if !bytes.Equal(data[:4], pktHeader) {
		err = ErrBadHeader
		return
	}
	body := make([]byte, len(data)-6)
	copy(body, data[:6])
	pkt = &CommPacket{
		Body:    body,
		Type:    CommType(data[4]),
		Version: ProtoVersion(data[5]),
	}
	return
}
