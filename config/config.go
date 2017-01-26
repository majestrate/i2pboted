package config

import (
	"github.com/majestrate/i2pboted/configparser"
	"github.com/majestrate/i2pboted/log"
)

type Config struct {
	I2P    I2PConfig
	Router RouterConfig
}

func Load(fname string) (*Config, error) {
	log.Debugf("load config from file: %s", fname)
	cfg, err := configparser.Read(fname)

	if err != nil {
		return nil, err
	}

	c := new(Config)

	s, _ := cfg.Section("i2p")
	c.I2P.load(s)
	s, _ = cfg.Section("bote")
	c.Router.load(s)

	return c, nil
}
