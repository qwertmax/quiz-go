package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qwertmax/quiz-go/cfg"
	// "github.com/qwertmax/quiz-go/db"
	"github.com/qwertmax/quiz-go/route"
)

func main() {
	config := cfg.Init("conf.yml")

	r := gin.Default()

	// var database db.Database
	// database.Init(config)

	// r.Use(db.DB(&database))

	//routings
	r.GET("/", route.Main)
	r.GET("/ping", route.Ping)

	r.Run(":" + config.APP_PORT)
}
