package network

import (
	"net"
)

type PeerInfo struct {
	// remote address of peer
	Addr net.Addr
}
