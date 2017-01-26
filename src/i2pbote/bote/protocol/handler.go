package protocol

import (
	"i2pbote/bote/protocol/comm"
	"net"
)

type Handler interface {
	// handle a communication packet
	CommPacket(pkt *comm.Packet, from net.Addr, c net.PacketConn) error
	// handle relay request
	RelayRequest(req *comm.RelayRequest, from net.Addr, c net.PacketConn) error
}

// create new packet handler
func NewHandler() Handler {
	return &handlerImpl{
		limiter: NewLimiter(),
	}
}
