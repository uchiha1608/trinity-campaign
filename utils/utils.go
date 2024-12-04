package utils

import (
	"trinity-campaign/config"
)

var (
	AppConfig config.AppConfig
)

func InitConfig(config config.AppConfig) {
	AppConfig = config
}
