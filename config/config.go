package config

import "github.com/spf13/viper"

type Config struct {
    
}

func LoadConfig(path string)(config *Config, err error){ 
    viper.AddConfigPath(path)
    viper.SetConfigName(".")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig(); 
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
