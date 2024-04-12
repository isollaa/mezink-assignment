package postgre

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func New(conf Config) *Connection {
	return &Connection{
		DB: conf.Connect(),
	}
}

func (p *Connection) RunMigration(conf Config) {
	log.Info().Msg("PostgreSQL Migrating...")
	db, _ := p.DB.DB()
	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: conf.DBName,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("An error occurred on postgres instance")
	}

	mgr, err := migrate.NewWithDatabaseInstance("file://./migration", "postgres", driver)
	if err != nil {
		log.Fatal().Err(err).Msg("An error occurred on initializing Migration instance")
	}

	err = mgr.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("An error occurred while syncing the database")
	}

	log.Info().Msg("PostgreSQL Migration Finished...")
}
