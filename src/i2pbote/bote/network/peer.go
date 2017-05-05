package network

import (
	"i2pbote/bote/protocol/comm"
	"net"
	"time"
)

type PeerInfo struct {
	// remote address of peer
	Addr net.Addr
}

type PeerSession interface {
	RecvPacket(pkt *comm.Packet) error
	SendPacket(pkt *comm.Packet) error
	LastContact() time.Time
	PeerInfo() PeerInfo
	TryConnect()
	// close session and remove from swarm
	Close()
}
