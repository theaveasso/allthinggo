package config

import "github.com/spf13/viper"

type Config struct {
	TG_BOT_TOKEN string `mapstructure:"TG_BOT_TOKEN"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
    viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
