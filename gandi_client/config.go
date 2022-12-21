package gandi_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	Key string `yaml:"key"  mapstructure:"key"`
}
