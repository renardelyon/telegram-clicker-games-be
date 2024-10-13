package config

import (
	gcfgv1 "gopkg.in/gcfg.v1"
)

func Setup() (*Config, error) {
	cfgFile := "config.ini"
	cfg := &Config{}
	err := gcfgv1.FatalOnly(gcfgv1.ReadFileInto(cfg, cfgFile))
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
