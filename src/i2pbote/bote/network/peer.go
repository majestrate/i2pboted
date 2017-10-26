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
	// handle inbound packet from this peer
	RecvPacket(pkt *comm.Packet) error
	// send a packet to remote peer
	SendPacket(pkt *comm.Packet) error
	// last time we had contact with this peer
	LastContact() time.Time
	// get remote peer's info
	PeerInfo() PeerInfo
	// try connecting one time
	TryConnect()
	// close session and remove from swarm
	Close()
}
