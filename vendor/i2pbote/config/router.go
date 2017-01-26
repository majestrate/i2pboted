package config

import (
	"i2pbote/configparser"
)

const DefaultDataDir = "i2pbote-data"

type RouterConfig struct {
	DataDir string
}

func (c *RouterConfig) load(s *configparser.Section) {
	if s == nil {
		c.DataDir = DefaultDataDir
	} else {
		c.DataDir = s.Get("datadir", DefaultDataDir)
	}
}
