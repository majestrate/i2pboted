package bote

import (
	"i2pbote/bote/network"
	"i2pbote/bote/protocol/comm"
	"i2pbote/config"
	"i2pbote/i2p"
	"i2pbote/log"
	"net"
)

type Router struct {
	session i2p.PacketSession
	done    chan error
	swarm   network.Swarm
	ready   bool
}

func NewRouter(cfg *config.RouterConfig) *Router {
	return &Router{
		done:  make(chan error),
		swarm: network.NewSwarm(cfg),
		ready: true,
	}
}

func (r *Router) InjectNetwork(s i2p.PacketSession) {
	log.Infof("Aquired net context with address %s", s.I2PAddr().Base32())
	r.session = s
	r.swarm.InjectNetwork(s)
}

func (r *Router) TryBootstrap(b network.Bootstrap) error {
	log.Infof("bootstrapping from %s", b.Name())
	peers, err := b.GetPeers()
	if err != nil {
		return err
	}
	for idx := range peers {
		r.swarm.EnsureSession(peers[idx])
	}
	return nil
}

func (r *Router) IsRunning() bool {
	return r.session != nil && r.ready
}

// close router and network context
func (r *Router) Close() {
	log.Info("closing i2pbote context...")
	if r.session != nil {
		log.Debug("closing network session")
		r.session.Close()
	}
	r.ready = false
	r.session = nil
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
	pkt, err := comm.Parse(data)
	if err != nil {
		log.Warnf("%s : %s", from, err)
		return
	}
	err = r.swarm.CommPacket(pkt, from)
	if err != nil {
		log.Warnf("packet error: %s", err)
		return
	}
}

// blocking run mainloop
func (r *Router) Run() {
	log.Debug("i2pbote run mainloop")
	for r.IsRunning() {
		var b [i2p.DatagramMTU]byte
		n, from, err := r.session.ReadFrom(b[:])
		if err != nil {
			r.done <- err
			return
		}
		go r.gotPacketFrom(b[:n], from)
	}
}
