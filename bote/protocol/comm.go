package protocol

import (
	"bytes"
	"errors"
	"github.com/majestrate/i2pboted/util"
)

// comm packet type byte
type CommType byte

// protocol version
type ProtoVersion byte

// i2pbote packet header
var pktHeader = []byte{0x6D, 0x30, 0x52, 0xe9}

var ErrTooSmall = errors.New("packet is too small")
var ErrBadHeader = errors.New("bad packet header")

const CommFetchReq = CommType(0x46)
const CommResponse = CommType(0x4e)
const CommPeerListReq = CommType(0x41)
const CommRelayReq = CommType(0x52)

// raw communication packet
type CommPacket struct {
	Type    CommType
	Version ProtoVersion
	Raw     []byte
}

func (pkt *CommPacket) Body() []byte {
	return pkt.Raw[6:]
}

// get as relay request
func (pkt *CommPacket) RelayRequest() (r *RelayRequest, err error) {
	if pkt.Type == CommRelayReq {
		body := pkt.Body()
		l := len(body)
		if l < (32 + 2 + 4 + 384 + 2) {
			err = ErrTooSmall
			return
		}
		r = new(RelayRequest)
		// correlation id
		copy(r.ID[:], body)
		// hashcash length
		hcl := util.UInt16_i(body[32:])
		if l < (32 + 2 + 4 + 384 + 2 + hcl) {
			err = ErrTooSmall
			return
		}
		// hash cash
		r.HashCash = make([]byte, hcl)
		copy(r.HashCash, body[32+2:])
		// delay
		r.Delay = util.UInt32(body[32+2+hcl:])
		// next hop
		copy(r.Next[:], body[32+2+4+hcl:])
		// encrypted data length
		dl := util.UInt16_i(body[32+2+4+384+hcl:])
		i := 32 + 3 + 4 + 384 + 2 + hcl
		// encrypted data
		r.Data = make([]byte, dl)
		copy(r.Data, body[i:i+dl])
		// rest is padding
	}
	return
}

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
	raw := make([]byte, len(data))
	copy(raw, data)
	pkt = &CommPacket{
		Raw:     raw,
		Type:    CommType(data[4]),
		Version: ProtoVersion(data[5]),
	}
	return
}
