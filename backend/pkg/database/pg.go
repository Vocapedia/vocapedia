package database

import (
	"fmt"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func pg(host string, port int, user, password, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, config.ReadValue().Database.SSLMode)
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			},
		),
	)
	return db, err
}
