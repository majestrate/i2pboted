package i2p

import (
	"crypto/sha256"
	"github.com/majestrate/i2pboted/i2p/base32"
	"github.com/majestrate/i2pboted/i2p/base64"
)

// base64 long form of i2p destination, implements net.Addr
type Addr string

// an i2p destination hash, the .b32.i2p address if you will
type DestHash [32]byte

// get string representation of i2p dest hash
func (h DestHash) String() string {
	b32addr := make([]byte, 56)
	base32.Encoding.Encode(b32addr, h[:])
	return string(b32addr[:52]) + ".b32.i2p"
}

// Returns the base64 representation of the I2PAddr
func (a Addr) Base64() string {
	return string(a)
}

// Returns the I2P destination (base64-encoded)
func (a Addr) String() string {
	return string(a)
}

// return base32 i2p desthash
func (a Addr) DestHash() (dh DestHash) {
	hash := sha256.New()
	b, _ := a.ToBytes()
	hash.Write(b)
	digest := hash.Sum(nil)
	copy(dh[:], digest)
	return
}

// return .b32.i2p address
func (a Addr) Base32() string {
	return a.DestHash().String()
}

// decode to i2p address to raw bytes
func (a Addr) ToBytes() (d []byte, err error) {
	buf := make([]byte, base64.Encoding.DecodedLen(len(a)))
	_, err = base64.Encoding.Decode(buf, []byte(a))
	if err == nil {
		d = buf
	}
	return
}

// Returns "I2P"
func (a Addr) Network() string {
	return "I2P"
}
