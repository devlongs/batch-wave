package config

import (
	"github.com/spf13/viper"
)

func InitConfig() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }
}

func GetConfig(key string) string {
    return viper.GetString(key)
}
