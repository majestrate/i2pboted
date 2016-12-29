package util

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"strings"
)

// generate a random string of length N
func RandStr(N int) (str string) {
	b := make([]byte, N)
	io.ReadFull(rand.Reader, b)
	str = base32.StdEncoding.EncodeToString(b)
	str = strings.ToLower(str[:N])
	return
}
