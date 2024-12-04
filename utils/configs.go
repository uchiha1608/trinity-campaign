package utils

import (
	"github.com/spf13/viper"
	"trinity-campaign/config"
)

func LoadConfig(path string) (conf config.AppConfig, err error) {
	viper.SetConfigFile(path + ".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)
	err = viper.Unmarshal(&conf.DBSetting)
	if err != nil {
		panic(err)
	}
	return
}
