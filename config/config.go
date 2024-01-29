package config

const (
	DefServiceHost = "0.0.0.0"
	DefServicePort = 9090
)

type (
	Config struct {
		Service *NodeConfig `mapstructure:"service" json:"service"`
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

func DefaultNodeConfig() *NodeConfig {
	return &NodeConfig{
		Host: DefServiceHost,
		Port: DefServicePort,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Service: DefaultNodeConfig(),
	}
}
