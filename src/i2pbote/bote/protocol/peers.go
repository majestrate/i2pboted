package protocol

import (
	"i2pbote/bote/common"
)

// peer storage
type PeerHolder interface {
	// get peers for relay
	GetPeers(limit int) []common.Destination
	AddPeers(d []common.Destination)
	Load() error
	Store() error
}
