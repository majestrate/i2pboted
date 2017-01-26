package comm

import (
	"bytes"
	"errors"
	"i2pbote/bote/common"
	"i2pbote/util"
)

// comm packet type byte
type PacketType byte

// protocol version
type ProtoVersion byte

// i2pbote packet header
var pktHeader = []byte{0x6D, 0x30, 0x52, 0xe9}

var ErrTooSmall = errors.New("packet is too small")
var ErrBadHeader = errors.New("bad packet header")

const FetchReq = PacketType(0x46)
const Response = PacketType(0x4e)
const PeerListReq = PacketType(0x41)
const RelayReq = PacketType(0x52)

// name of comm packet type
func (t PacketType) Name() string {
	switch t {
	case RelayReq:
		return "RelayRequest"
	case PeerListReq:
		return "PeerListRequest"
	case Response:
		return "Response"
	case FetchReq:
		return "FetchRequest"
	default:
		return "Unknown"
	}
}

// raw communication packet
type Packet struct {
	Type    PacketType
	Version ProtoVersion
	Raw     []byte
}

func (pkt *Packet) Body() []byte {
	return pkt.Raw[6:]
}

// get as peer list packet
func (pkt *Packet) PeerList() (cid common.CID, err error) {
	if pkt.Type == PeerListReq {
	}
	return
}

// get as relay request
func (pkt *Packet) RelayRequest() (r *RelayRequest, err error) {
	if pkt.Type == RelayReq {
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
func Parse(data []byte) (pkt *Packet, err error) {
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
	pkt = &Packet{
		Raw:     raw,
		Type:    PacketType(data[4]),
		Version: ProtoVersion(data[5]),
	}
	return
}
