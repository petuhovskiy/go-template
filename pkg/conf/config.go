package conf

import (
	"github.com/caarlos0/env/v6"
)

type Struct struct {
	PrometheusBind string `env:"PROMETHEUS_BIND" envDefault:":2112"`
}

func ParseEnv() (*Struct, error) {
	cfg := Struct{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
