package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	DefaultConfigFilePath = "./config.yml"
	DefServiceHost        = "0.0.0.0"
	DefServicePort        = 9090
)

type (
	Config struct {
		Service     *NodeConfig `mapstructure:"service" json:"service"`
		ContentPath string      `mapstructure:"content_path" json:"content_path"`
	}
	NodeConfig struct {
		Host string `mapstructure:"host" json:"host"`
		Port int    `mapstructure:"port" json:"port"`
	}
)

var AppConfig *Config

func init() {
	AppConfig = DefaultConfig()

}

func LoadFile(filepath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(filepath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("configuration file not found")
		} else {
			return nil, errors.Wrap(err, "ReadInConfig")
		}
	}
	config := DefaultConfig()
	if err := v.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}
	AppConfig = config
	return config, nil
}

func DefaultNodeConfig() *NodeConfig {
	return &NodeConfig{
		Host: DefServiceHost,
		Port: DefServicePort,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Service:     DefaultNodeConfig(),
		ContentPath: "content",
	}
}
