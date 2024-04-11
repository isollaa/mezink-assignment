package postgre

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	maxIdleConnection = 10
	maxOpenConnection = 10
	maxLifeTime       = 5 * time.Minute
)

type Config struct {
	DBName          string
	Host            string
	Port            int
	Username        string
	Password        string
	Timezone        string
	Debug           bool
	EnableMigration bool
}

func (c *Config) Connect() *gorm.DB {
	descriptor := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		c.Host,
		c.Username,
		c.Password,
		c.DBName,
		c.Port,
		url.QueryEscape(c.Timezone))

	sqlDB, err := sql.Open("pgx", descriptor)
	if err != nil {
		log.Err(err).
			Str("username", c.Username).
			Str("host", c.Host).
			Int("port", c.Port).
			Str("dbName", c.DBName).
			Msg("[postgre] failed to connect to database")
		panic("[postgre] failed to open sql db")
	}

	log.Info().
		Str("username", c.Username).
		Str("host", c.Host).
		Int("port", c.Port).
		Str("dbName", c.DBName).
		Msg("[postgre] connected to database")

	sqlDB.SetMaxIdleConns(maxIdleConnection)
	sqlDB.SetMaxOpenConns(maxOpenConnection)
	sqlDB.SetConnMaxLifetime(maxLifeTime)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		panic("[postgre] failed to initialize gorm with sql db")
	}

	//log query on debug = true
	if c.Debug {
		return gormDB.Debug()
	}

	return gormDB
}
