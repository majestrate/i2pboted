package network

// network bootstrapping mechanism
type Bootstrap interface {
	// get some peers from a bootstrap server or peer
	GetPeers() ([]*PeerInfo, error)
}
