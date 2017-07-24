package config

import (
	"i2pbote/configparser"
	"i2pbote/log"
	"i2pbote/util"
)

type Config struct {
	Bootstrap BootstrapConfig
	I2P       I2PConfig
	Router    RouterConfig
	RPC       RPCConfig
}

type configLoadable interface {
	Load(sect *configparser.Section)
}

func Load(fname string) (*Config, error) {
	err := util.EnsureFile(fname, 0)
	if err != nil {
		return nil, err
	}
	log.Debugf("load config from file: %s", fname)
	cfg, err := configparser.Read(fname)

	if err != nil {
		return nil, err
	}

	c := new(Config)

	m := map[string]configLoadable{
		"i2p":       &c.I2P,
		"bootstrap": &c.Bootstrap,
		"bote":      &c.Router,
		"rpc":       &c.RPC,
	}

	for k, v := range m {
		s, _ := cfg.Section(k)
		v.Load(s)
	}

	return c, nil
}
