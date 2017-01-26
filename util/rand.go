package util

import (
	"crypto/rand"
	"github.com/majestrate/i2pboted/i2p/base32"
	"io"
	"strings"
)

// generate a random string of length N
func RandStr(N int) (str string) {
	b := make([]byte, N)
	io.ReadFull(rand.Reader, b)
	str = base32.Encoding.EncodeToString(b)
	str = strings.ToLower(str[:N])
	return
}
