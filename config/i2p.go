package config

import (
	"github.com/majestrate/i2pboted/configparser"
)

const DefaultI2PAddr = "127.0.0.1:7656"

type I2PConfig struct {
	Addr        string
	Keyfile     string
	SessionName string
}

func (c I2PConfig) load(s *configparser.Section) {
	if s == nil {
		c.Addr = DefaultI2PAddr
	} else {
		c.Addr = s.Get("addr", DefaultI2PAddr)
		c.Keyfile = s.Get("keyfile", "")
		c.SessionName = s.Get("session", "")
	}
}
