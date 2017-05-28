package app

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/metaops/metaops-api/config"
	"log"
	"os"
)

type App struct {
	Logger *log.Logger

	config *config.Config
	db     *gorm.DB
	redis  redis.Conn
}

func New(appConfig *config.Config) *App {

	return &App{
		Logger: log.New(os.Stdout, "metaops-api: ", log.Lshortfile|log.LstdFlags),
		db:     initDB(appConfig.DBConfig),
		redis:  initRedis(appConfig.RedisConfig),
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

func initRedis(redisConfig config.RedisConfig) redis.Conn {
	conn, err := redis.DialURL(redisConfig.URL)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
