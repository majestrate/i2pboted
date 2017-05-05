package config

import (
	"i2pbote/configparser"
)

const DefaultBootstrapAddr = "f2q3h7ehh73wo3gbzbsidqhoalyq5djfcsk3cuglp56rd3ibudua.b32.i2p"
const DefaultBootstrapFile = "nodes.txt"

type BootstrapConfig struct {
	NodeAddr string
	NodeFile string
}

func (c *BootstrapConfig) Load(s *configparser.Section) {
	if s == nil {
		c.NodeFile = DefaultBootstrapFile
		c.NodeAddr = DefaultBootstrapAddr
	} else {
		c.NodeFile = s.Get("file", DefaultBootstrapFile)
		c.NodeAddr = s.Get("addr", DefaultBootstrapAddr)
	}
}
