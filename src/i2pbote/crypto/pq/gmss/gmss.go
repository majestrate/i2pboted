package gmss

import (
	"crypto"
	"math"
)

type privKey struct {
}

type winternitzOTSignature struct {
	privkeySize uint32
	digest      crypto.Hash
	w           uint32
}

func createWinternitzOTSignature(w uint32, d crypto.Hash) *winternitzOTSignature {
	mdbits := d.Size() << 3
	ksize := math.Ceil(float64(mdbits))
	return &winternitzOTSignature{
		w:           w,
		digest:      d,
		privkeySize: uint32(ksize),
	}
}
