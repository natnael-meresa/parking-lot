package config

import "github.com/spf13/viper"

// Config is the struct that holds the application configuration.
type Config struct {
	Server ServerConfig `json:"server"`
	DB     DBConfig     `json:"db"`
}
// ServerConfig is the struct that holds the server configuration.
type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
// DBConfig is the struct that holds the database configuration.
type DBConfig struct {
	URL     string `json:"url"`
	PoolMax string `json:"pool_max"`
}

// SetDefaults sets the default values for the server configuration.
func SetDefaults() {
	viper.SetDefault("SERVER_HOST", "localhost")
	viper.SetDefault("SERVER_PORT", 8000)
}

// NewConfig creates a new instance of the application configuration.
func NewConfig() (*Config, error) {
	config := &Config{}

	// Set the default values for the server configuration.
	SetDefaults()
	viper.SetConfigType("yml")
	viper.SetConfigFile("./config/config.yaml")
	viper.AutomaticEnv()

	// Read the configuration file.
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Unmarshal the configuration file into the Config struct.
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
