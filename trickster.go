package main

import (
	"github.com/BurntSushi/toml"
)

type tricksterConfig struct {
	Origins map[string]prometheusOriginConfig `toml:"origins"`
}

type prometheusOriginConfig struct {
	OriginURL           string `toml:"origin_url"`
	APIPath             string `toml:"api_path"`
	IgnoreNoCacheHeader bool   `toml:"ignore_no_cache_header"`
	MaxValueAgeSecs     int64  `toml:"max_value_age_secs"`
	FastForwardDisable  bool   `toml:"fast_forward_disable"`
}

func (c *tricksterConfig) loadFile(path string) error {
	_, err := toml.DecodeFile(path, &c)
	return err
}
