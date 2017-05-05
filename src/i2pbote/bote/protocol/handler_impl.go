package protocol

import (
	"i2pbote/bote/common"
	"i2pbote/bote/protocol/comm"
	"i2pbote/bote/protocol/data"
	"i2pbote/log"
	"net"
)

type handlerImpl struct {
	limiter     *Limiter
	relay_peers PeerHolder
	conn        net.PacketConn
}

// handle comm packet
func (h *handlerImpl) CommPacket(pkt *comm.Packet, from net.Addr) (err error) {

	if h.limiter.CheckRX(len(pkt.Raw), from) {
		// drop too high rx
		log.Warnf("RX too high, dropping packet from %s", from)
		return
	}
	switch pkt.Type() {
	case comm.RelayReq:
		req, e := pkt.RelayRequest()
		if e != nil {
			err = e
			return
		}
		return h.RelayRequest(req, from)
	case comm.PeerListReq:
		req, e := pkt.PeerListRequest()
		if e != nil {
			err = e
			return
		}
		return h.SendPeerList(req.CID, from)
	}
	return
}

func (h *handlerImpl) RelayRequest(req *comm.RelayRequest, from net.Addr) (err error) {

	return
}

// send current peer list to a remote
func (h *handlerImpl) SendPeerList(cid common.CID, to net.Addr) (err error) {
	peers := h.relay_peers.GetPeers(20)
	buff := data.CreatePeerList(Version.Byte(), peers)
	pkt := comm.ResponsePacket(Version.Byte(), comm.OK, cid, buff)

	_, err = h.conn.WriteTo(pkt.Raw[:], to)
	return
}
