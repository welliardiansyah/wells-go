package config

import (
	"sync"

	"github.com/spf13/viper"
)

var configInstance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config, err := LoadConfig(".")
		if err != nil {
			panic("Error loading config: " + err.Error())
		}
		configInstance = &config
	})
	return configInstance
}

type Config struct {
	Environment          string `mapstructure:"ENVIRONMENT"`
	BaseURL              string `mapstructure:"BASE_URL"`
	AllowedOrigins       string `mapstructure:"ALLOWED_ORIGINS"`
	AllowedMethods       string `mapstructure:"ALLOWED_METHODS"`
	AuthorizationHeaders string `mapstructure:"AUTHORIZATION_HEADERS"`
	DbUser               string `mapstructure:"DB_USER"`
	DbPassword           string `mapstructure:"DB_PASSWORD"`
	DbHost               string `mapstructure:"DB_HOST"`
	DbPort               string `mapstructure:"DB_PORT"`
	DbName               string `mapstructure:"DB_NAME"`
	ServerPort           string `mapstructure:"SERVER_PORT"`
	DBSource             string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress    string `mapstructure:"HTTP_SERVER_ADDRESS"`

	// 🔑 JWT Config
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	JWTIssuer       string `mapstructure:"JWT_ISSUER"`
	AccessTokenTTL  string `mapstructure:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL string `mapstructure:"REFRESH_TOKEN_TTL"`
}

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
