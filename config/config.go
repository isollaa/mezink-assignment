package config

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	InfraPosgreDBName          string `mapstructure:"INFRA_POSGRE_DB_NAME"`
	InfraPosgreEnableDebug     bool   `mapstructure:"INFRA_POSGRE_ENABLE_DEBUG"`
	InfraPosgreEnableMigration bool   `mapstructure:"INFRA_POSGRE_ENABLE_MIGRATION"`
	InfraPosgreHost            string `mapstructure:"INFRA_POSGRE_HOST"`
	InfraPosgrePassword        string `mapstructure:"INFRA_POSGRE_PASSWORD"`
	InfraPosgrePort            int    `mapstructure:"INFRA_POSGRE_PORT"`
	InfraPosgreTimezone        string `mapstructure:"INFRA_POSGRE_TIMEZONE"`
	InfraPosgreUsername        string `mapstructure:"INFRA_POSGRE_USERNAME"`
	ServerName                 string `mapstructure:"SERVER_NAME"`
	ServerPort                 int    `mapstructure:"SERVER_PORT"`
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data an return the struct
func Get() Config {
	once.Do(func() {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read config file")
		}

		log.Info().Msg("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to unmarshal config file")
		}
	})

	return conf
}
