package config

import (
	"i2pbote/configparser"
)

const DefaultBootstrapFile = "nodes.txt"

type BootstrapConfig struct {
	NodeFile string
}

func (c *BootstrapConfig) Load(s *configparser.Section) {
	if s == nil {
		c.NodeFile = DefaultBootstrapFile
	} else {
		c.NodeFile = s.Get("file", DefaultBootstrapFile)
	}
}
