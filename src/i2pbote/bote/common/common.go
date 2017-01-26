package common

import (
	"i2pbote/i2p"
	"i2pbote/i2p/base64"
)

// correlation id
type CID [32]byte

func (cid CID) String() string {
	return base64.Encoding.EncodeToString(cid[:])
}

// i2p destination
type Destination [384]byte

// base64 string representation
func (d Destination) String() string {
	return base64.Encoding.EncodeToString(d[:])
}

// convert to i2p address blob
func (d Destination) ToAddr() i2p.Addr {
	return i2p.Addr(d.String())
}
