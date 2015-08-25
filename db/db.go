package db

import (
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qwertmax/quiz-go/cfg"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) getDb(conf cfg.Config) (gorm.DB, error) {
	dbconn := "user=" + conf.DB_USERNAME + " password=" + conf.DB_PASSWORD + " dbname=" + conf.DB_NAME + " sslmode=" + conf.DB_SSLMODE + " host=" + conf.DB_ADDRESS + " port=" + conf.DB_PORT
	return gorm.Open("postgres", dbconn)
}

func (db *Database) Init(conf cfg.Config) {
	dbHandler, err := db.getDb(conf)
	if err != nil {
		fmt.Println(dbHandler)
	}

	err = dbHandler.DB().Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	db.DB = &dbHandler

	dbHandler.DB()
	dbHandler.DB().SetMaxIdleConns(10)
	dbHandler.DB().SetMaxOpenConns(100)
	dbHandler.SingularTable(true)

	// dbHandler.LogMode(true)
}

// Connect to DB and return DB connection handler
func DB(db *Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
