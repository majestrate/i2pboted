package bote

import (
	"i2pbote/bote/protocol"
	"i2pbote/config"
	"i2pbote/i2p"
	"i2pbote/log"
	"net"
)

type Router struct {
	session i2p.PacketSession
	done    chan error
}

func NewRouter(cfg config.RouterConfig) *Router {
	return &Router{
		done: make(chan error),
	}
}

func (r *Router) InjectNetwork(s i2p.PacketSession) {
	log.Infof("Aquired net context with address %s", s.I2PAddr().Base32())
	r.session = s
}

// close router and network context
func (r *Router) Close() {
	log.Info("closing i2pbote context...")
	if r.session != nil {
		log.Debug("closing network session")
		r.session.Close()
	}
	r.done <- nil
}

func (r *Router) GracefulClose() {
	log.Info("graceful close started")
	// TODO: begin graceful shutdown
}

// get bote rpc instance
func (r *Router) RPC() *RPC {
	return &RPC{
		router: r,
	}
}

// wait until done
func (r *Router) Wait() error {
	err := <-r.done
	close(r.done)
	return err
}

func (r *Router) gotPacketFrom(data []byte, from net.Addr) {
	log.Debugf("got %d bytes from remote peer", len(data))
	pkt, err := protocol.ParseCommPacket(data)
	if err != nil {
		log.Warnf("%s : %s", from, err)
		return
	}
	r.handleCommPacketFrom(pkt, from)
}

func (r *Router) handleCommPacketFrom(pkt *protocol.CommPacket, from net.Addr) {

	log.Debugf("got CommPacket %s", pkt.Type.Name())
	switch pkt.Type {
	case protocol.CommRelayReq:
		relayReq, err := pkt.RelayRequest()
		if err == nil {
			r.handleRelayRequest(relayReq, from)
		} else {
			log.Errorf("bad relay request packet: %s", err)
		}
	}
}

func (r *Router) handleRelayRequest(req *protocol.RelayRequest, from net.Addr) {
	nextAddr := req.Next.ToAddr()
	log.Debugf("got relay request to %s", nextAddr.Base32())
}

// blocking run mainloop
func (r *Router) Run() {
	log.Debug("i2pbote run mainloop")
	var b [i2p.DatagramMTU]byte
	for {
		n, from, err := r.session.ReadFrom(b[:])
		if err != nil {
			r.done <- err
			return
		}
		msg := make([]byte, n)
		copy(msg, b[:n])
		go r.gotPacketFrom(msg, from)
	}
}
