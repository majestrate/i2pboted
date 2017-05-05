package network

import (
	"i2pbote/bote/protocol"
	"i2pbote/bote/protocol/comm"
	"i2pbote/bote/protocol/data"
	"i2pbote/log"
	"net"
	"time"
)

type peerSessionImpl struct {
	addr   net.Addr
	parent *swarmImpl
	conn   net.PacketConn
	lastRX time.Time
}

func (p *peerSessionImpl) sendClose() {

}

func (p *peerSessionImpl) Close() {
	parent := p.parent
	parent.access.Lock()
	if _, ok := parent.peers[p.addr]; ok {
		p.sendClose()
		delete(parent.peers, p.addr)
	}
	parent.access.Unlock()
}

func (p *peerSessionImpl) RecvPacket(pkt *comm.Packet) (err error) {
	p.lastRX = time.Now()
	t := pkt.Type()
	switch t {
	case comm.RelayReq:
		_, e := pkt.RelayRequest()
		err = e
		break
	case comm.PeerListReq:
		req, e := pkt.PeerListRequest()
		err = e
		if err == nil {
			peers := p.parent.getPeers(20)
			buff := data.CreatePeerList(protocol.Version.Byte(), peers)
			err = p.SendPacket(comm.ResponsePacket(protocol.Version.Byte(), comm.OK, req.CID, buff))
		}
		break
	default:
		log.Warnf("got weird comm packet of type %s", t)
	}
	return
}

func (p *peerSessionImpl) SendPacket(pkt *comm.Packet) (err error) {
	_, err = p.conn.WriteTo(pkt.Raw[:], p.addr)
	return
}

func (p *peerSessionImpl) LastContact() time.Time {
	return p.lastRX
}

func (p *peerSessionImpl) PeerInfo() (i PeerInfo) {
	i.Addr = p.addr
	return
}

func (p *peerSessionImpl) TryConnect() {

}
