package config

import "github.com/jinzhu/configor"

type Config struct {
	SelfURL string `json:"self_url" default:"https://nas.local"`
	SelfPort int `json:"self_port" default:"8888" required:"true"`
	TransmissionConfig Transmission `json:"transmission_config"`
}

type Transmission struct {
	PublicURL string `json:"public_url" default:"http://nas.local:9091" env:"CONFIG__TRANSMISSION__PUBLIC_URL" required:"true"`
	Endpoint string `json:"endpoint" default:"http://localhost:9091" env:"CONFIG__TRANSMISSION__ENDPOINT" required:"true"`
}

func MustLoadConfig() *Config {
	cfg := &Config{}
	err := configor.New(&configor.Config{AutoReload: false, Silent: true}).Load(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
