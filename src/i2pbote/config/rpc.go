package config

import (
	"errors"
	"i2pbote/configparser"
	"net"
)

type RPCConfig struct {
	Enabled  bool
	BindUnix string
	BindInet string
}

func (c *RPCConfig) Load(s *configparser.Section) {
	if s != nil {
		c.Enabled = s.Get("enable", "0") == "1"
		c.BindInet = s.Get("inet", "")
		c.BindUnix = s.Get("unix", "botectl.sock")
	}
}

// create net.Listener for rpc
func (c *RPCConfig) CreateListener() (net.Listener, error) {
	if c.Enabled {
		if c.BindUnix != "" {
			return net.Listen("unix", c.BindUnix)
		}
		if c.BindInet != "" {
			return net.Listen("tcp", c.BindInet)
		}
		return nil, errors.New("no rpc address specified")
	}
	return nil, errors.New("rpc not enabled")
}
