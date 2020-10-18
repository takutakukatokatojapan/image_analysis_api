package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	BindPort string `envconfig:"BIND_PORT" default:":8080"`
}

// NewConfig
//	Config構造体を生成する
func NewConfig() Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
	}
	return config
}
