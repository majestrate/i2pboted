package protocol

import (
	"i2pbote/bote/protocol/comm"
	"i2pbote/config"
	"net"
)

type Handler interface {
	// handle a communication packet
	CommPacket(pkt *comm.Packet, from net.Addr, c net.PacketConn) error
	// handle relay request
	RelayRequest(req *comm.RelayRequest, from net.Addr, c net.PacketConn) error
}

// create new packet handler
func NewHandler(c *config.RouterConfig) Handler {
	return &handlerImpl{
		limiter:     NewLimiter(),
		relay_peers: NewFsPeerHolder(c.DataDir),
	}
}
