package infra

import (
	"mezink-assignment/config"
	"mezink-assignment/infra/postgre"
)

type Connection struct {
	PG *postgre.Connection
}

func NewConnection(conf Config) *Connection {
	conn := Connection{
		PG: postgre.New(conf.PG),
	}

	if conf.PG.EnableMigration {
		conn.PG.RunMigration(conf.PG)
	}

	return &conn
}

type Config struct {
	PG postgre.Config
}

func NewConfig(conf config.Config) Config {
	return Config{
		PG: postgre.Config{
			DBName:          conf.InfraPosgreDBName,
			Host:            conf.InfraPosgreHost,
			Password:        conf.InfraPosgrePassword,
			Port:            conf.InfraPosgrePort,
			Timezone:        conf.InfraPosgreTimezone,
			Username:        conf.InfraPosgreUsername,
			Debug:           conf.InfraPosgreEnableDebug,
			EnableMigration: conf.InfraPosgreEnableMigration,
		},
	}
}
