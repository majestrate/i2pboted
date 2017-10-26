package gmss

import (
	"crypto"
	"io"
)

type GMSSSigner struct {
	priv *PrivateKey
}

func (s *GMSSSigner) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (sig []byte, err error) {
	return
}
