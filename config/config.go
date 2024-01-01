package config

import "github.com/spf13/viper"

type Config struct {
	Server ServerConfig `json:"server"`
	DB     DBConfig     `json:"db"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DBConfig struct {
	URL     string `json:"url"`
	PoolMax string `json:"pool_max"`
}

func SetDefaults() {
	viper.SetDefault("SERVER_HOST", "localhost")
	viper.SetDefault("SERVER_PORT", 8000)
}

func NewConfig() (*Config, error) {
	config := &Config{}

	SetDefaults()
	viper.SetConfigType("yml")
	viper.SetConfigFile("./config/config.yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
