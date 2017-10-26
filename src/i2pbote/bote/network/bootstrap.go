package network

import (
	"bufio"
	"bytes"
	"i2pbote/i2p"
	"io/ioutil"
	"strings"
)

// network bootstrapping mechanism
type Bootstrap interface {
	// get some peers from a bootstrap server or peer
	GetPeers() ([]*PeerInfo, error)
	// name of bootstrap method
	Name() string
}

// bootstrap from file
type FileBootstrap struct {
	name    string
	session i2p.Session
}

/*
func (a *FileBootstrap) GetPeers() (peers []*PeerInfo, err error) {
	var addr net.Addr
	addr, err = a.session.LookupI2P(a.name)
	if err == nil {
		peers = append(peers, &PeerInfo{
			Addr: addr,
		})
	}
	return
}
*/

func (a *FileBootstrap) GetPeers() (peers []*PeerInfo, err error) {
	var data []byte
	data, err = ioutil.ReadFile(a.name)
	if err == nil {
		r := bytes.NewReader(data)
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			line := sc.Text()
			idx := strings.Index(line, "#")
			if idx >= 0 {
				// comment
				// TODO: trim
				continue
			}
			peers = append(peers, &PeerInfo{
				Addr: i2p.Addr(line),
			})
		}
	}
	return
}

func (a *FileBootstrap) Name() string {
	return "file bootstrap: " + a.name
}

func NewFileBootstrap(name string, s i2p.Session) *FileBootstrap {
	return &FileBootstrap{
		name:    name,
		session: s,
	}
}
