package comm

import (
	"i2pbote/bote/common"
	"i2pbote/bote/protocol/chain"
)

type RelayRequest struct {
	ID       common.CID
	Delay    uint32
	Next     common.Destination
	Return   chain.ReturnChain
	Data     []byte
	HashCash []byte
}
