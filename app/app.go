package app

import (
	"github.com/jinzhu/gorm"
	"github.com/metaops/metaops-api/config"
	"log"
	"os"
)

type App struct {
	Logger *log.Logger

	config *config.Config
	db     *gorm.DB
}

func New(appConfig *config.Config) *App {

	return &App{
		Logger: log.New(os.Stdout, "metaops-api: ", log.Lshortfile|log.LstdFlags),
		db:     initDB(appConfig.DBConfig),
	}
}

func initDB(dbConfig config.DBConfig) *gorm.DB {
	db, err := gorm.Open(dbConfig.Driver, dbConfig.Connection)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	return db
}
