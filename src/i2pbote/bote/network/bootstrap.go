package network

import (
	"i2pbote/i2p"
	"net"
)

// network bootstrapping mechanism
type Bootstrap interface {
	// get some peers from a bootstrap server or peer
	GetPeers() ([]*PeerInfo, error)
	// name of bootstrap method
	Name() string
}

// bootstrap from 1 peer
type NameBootstrap struct {
	name    string
	session i2p.Session
}

// get 1 peer , the node itself
func (a *NameBootstrap) GetPeers() (peers []*PeerInfo, err error) {
	var addr net.Addr
	addr, err = a.session.LookupI2P(a.name)
	if err == nil {
		peers = append(peers, &PeerInfo{
			Addr: addr,
		})
	}
	return
}

func (a *NameBootstrap) Name() string {
	return "Seed Node: " + a.name
}

func NewNameBootstrap(name string, s i2p.Session) *NameBootstrap {
	return &NameBootstrap{
		name:    name,
		session: s,
	}
}
