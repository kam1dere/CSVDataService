package config

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres struct {
		URL    string `json:"url"`
		Driver string `json:"driver"`
	} `json:"postgres"`

	Server struct {
		Network    string `json:"network"`
		Address    string `json:"address"`
		RestServer string `json:"restServer"`
	} `json:"server"`

	Logger struct {
		ServiceName string `json:"serviceName"`
	} `json:"logger"`

	System struct {
		DeviceIdForLiquid string `json:"deviceIdForLiquid"`
		Services          struct {
			Auth struct {
				Url string `json:"url"`
			} `json:"auth"`
		} `json:"system"`
	}
}

// LoadConfig loads config from specified --config parameter
func LoadConfig(configFile string) (*viper.Viper, error) {
	vconfig := viper.New()

	vconfig.SetConfigFile(configFile)
	vconfig.AutomaticEnv()
	vconfig.SetEnvPrefix("SERVICE")
	vconfig.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := vconfig.ReadInConfig()
	if err != nil {
		log.Err(err).Stack().Msg("")

		return nil, fmt.Errorf("read config error: %w", err)
	}

	return vconfig, nil
}

// ParseConfig parses the configuration from YAM into struct.
func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarshall config error: %w", err)
	}

	return &config, nil
}
