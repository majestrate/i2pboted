package protocol

import (
	"i2pbote/bote/protocol/comm"
	"i2pbote/log"
	"net"
)

type handlerImpl struct {
	limiter *Limiter
}

// handle comm packet
func (h *handlerImpl) CommPacket(pkt *comm.Packet, from net.Addr, c net.PacketConn) (err error) {

	if h.limiter.CheckRX(len(pkt.Raw), from) {
		// drop too high rx
		log.Warnf("RX too high, dropping packet from %s", from)
		return
	}
	switch pkt.Type {
	case comm.RelayReq:
		req, e := pkt.RelayRequest()
		if e != nil {
			err = e
			return
		}
		return h.RelayRequest(req, from, c)
	}
	return
}

func (h *handlerImpl) RelayRequest(req *comm.RelayRequest, from net.Addr, c net.PacketConn) (err error) {

	return
}
