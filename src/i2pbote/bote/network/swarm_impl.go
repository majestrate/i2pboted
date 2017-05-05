package network

import (
	"i2pbote/bote/common"
	"i2pbote/bote/protocol/comm"
	"i2pbote/config"
	"net"
	"sync"
)

type swarmImpl struct {
	conf   *config.RouterConfig
	peers  map[net.Addr]PeerSession
	access sync.Mutex
	conn   net.PacketConn
}

func (sw *swarmImpl) EnsureSession(info *PeerInfo) (session PeerSession) {
	sw.access.Lock()
	var ok bool
	session, ok = sw.peers[info.Addr]
	if !ok {
		session = &peerSessionImpl{
			addr:   info.Addr,
			parent: sw,
			conn:   sw.conn,
		}
		sw.peers[info.Addr] = session
	}
	sw.access.Unlock()
	return
}

func (sw *swarmImpl) ForEachSession(v func(PeerSession)) {
	var sessions []PeerSession
	sw.access.Lock()
	for k := range sw.peers {
		sessions = append(sessions, sw.peers[k])
	}
	sw.access.Unlock()
	for idx := range sessions {
		v(sessions[idx])
	}
}

func (sw *swarmImpl) getPeers(num int) (peers []common.Destination) {
	sw.ForEachSession(func(s PeerSession) {
		if num > 0 {
			peers = append(peers, common.AddrToDest(s.PeerInfo().Addr))
		}
		num--
	})
	return
}

func (sw *swarmImpl) InjectNetwork(c net.PacketConn) {
	sw.conn = c
}

func (sw *swarmImpl) CommPacket(pkt *comm.Packet, from net.Addr) error {
	sess := sw.EnsureSession(&PeerInfo{
		Addr: from,
	})
	return sess.RecvPacket(pkt)
}
