package network

import (
	"i2pbote/bote/protocol/comm"
	"i2pbote/config"
	"net"
)

// Swarm is network state for all peers we are talking with
type Swarm interface {
	// handle inbound comm packet
	CommPacket(*comm.Packet, net.Addr) error
	// EnsurePeer adds peer session if we don't have a session with a peer or gets an existing one if it's not there
	EnsureSession(info *PeerInfo) PeerSession
	// visit each session we have in this swarm
	ForEachSession(v func(PeerSession))
	// inject udp socket
	InjectNetwork(net.PacketConn)
}

func NewSwarm(c *config.RouterConfig) Swarm {
	return &swarmImpl{
		conf:  c,
		peers: make(map[net.Addr]PeerSession),
	}
}
