package protocol

import (
	"i2pbote/bote/protocol/comm"
	"net"
)

type Handler interface {
	CommPacket(pkt *comm.Packet, from net.Addr) (*comm.Packet, error)
}
