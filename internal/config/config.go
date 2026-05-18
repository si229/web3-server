package config

import "github.com/spf13/viper"

type Route struct {
	Method  string `mapstructure:"method"`
	Path    string `mapstructure:"path"`
	Handler string `mapstructure:"handler"`
}

type Config struct {
	Routes []Route `mapstructure:"routes"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile("config/routes.yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	err := v.Unmarshal(&cfg)
	return &cfg, err
}
