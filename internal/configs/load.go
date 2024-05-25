package configs

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	ErrConfigLoad = errors.New("error loading configs")
)

func LoadConfigs() (*Configs, error) {
	log.Info("Loading configurations.")
	var (
		config   Configs
		provided bool
	)

	if config.Endpoint, provided = os.LookupEnv("SERVER_ENDPOINT"); !provided {
		return nil, ErrConfigLoad
	}
	if config.KrakenEndpoint, provided = os.LookupEnv("KRAKEN_ENDPOINT"); !provided {
		return nil, ErrConfigLoad
	}

	return &config, nil
}
