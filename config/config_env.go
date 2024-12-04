package config

type AppConfig struct {
	DBSetting struct {
		Port           string `mapstructure:"DB_PORT"`
		Name           string `mapstructure:"DB_NAME"`
		MaxConnections int    `mapstructure:"DB_MAX_CONNECTIONS"`
	}
	HttpsConfig struct {
		AllowedOrigins   []string `mapstructure:"ALLOWED_ORIGINS"`
		AllowedMethods   []string `mapstructure:"ALLOWED_METHODS"`
		AllowedHeaders   []string `mapstructure:"ALLOWED_HEADERS"`
		AllowCredentials bool     `mapstructure:"ALLOW_CREDENTIALS"`
		MaxAge           int      `mapstructure:"MAX_AGE"`
	}
}
