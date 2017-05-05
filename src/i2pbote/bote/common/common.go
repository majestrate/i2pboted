package common

import (
	"i2pbote/i2p"
	"i2pbote/i2p/base64"
	"net"
)

const CIDLen = 32

// correlation id
type CID [CIDLen]byte

func (cid CID) String() string {
	return base64.Encoding.EncodeToString(cid[:])
}

const DestLen = 384

// i2p destination
type Destination [DestLen]byte

// base64 string representation
func (d Destination) String() string {
	return base64.Encoding.EncodeToString(d[:])
}

// convert to i2p address blob
func (d Destination) ToAddr() i2p.Addr {
	return i2p.Addr(d.String())
}

func AddrToDest(a net.Addr) (d Destination) {
	i := i2p.Addr(a.String())
	b, e := i.ToBytes()
	if e == nil {
		copy(d[:], b[:DestLen])
	}
	return
}
