package database

import (
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDB(host string, port int, user, password, dbname string) {
	db, err = pg(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.User{})
}
func Manager() *gorm.DB {
	return db
}
