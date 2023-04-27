package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config contains the configuration for the server
// the values are read by viper from the config file or environment variables

type Config struct {
	DBDriver            string        `mapstructure:"db_driver"`
	DBSource            string        `mapstructure:"db_source"`
	ServerAddress       string        `mapstructure:"server_address"`
	TokenSymmetricKey   string        `mapstructure:"token_symmetric_key"`
	AccessTokenDuration time.Duration `mapstructure:"access_token_duration"`
}

// LoadConfig loads the configuration from the config file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
