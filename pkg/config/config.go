package config

import "github.com/spf13/viper"

type Config struct {
	User     string `mapstructure:"DNS_O_MATIC_USER"`
	Password string `mapstructure:"DNS_O_MATIC_PASSWORD"`
	LogFile  string `mapstructure:"DNS_O_MATIC_LOG_FILE"`
}

// LoadConfig gets the config necessary from app.env and overrides it with ENV variables
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
	return
}
