package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qwertmax/quiz-go/cfg"
	// "github.com/qwertmax/quiz-go/db"
	"github.com/qwertmax/quiz-go/route"
	"net/http"
)

func main() {
	config := cfg.Init("conf.yml")
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// r.Static("/assets", "./assets")
	r.StaticFS("/css", http.Dir("res/js"))
	r.StaticFS("/js", http.Dir("res/js"))
	r.StaticFile("/favicon.ico", "res/favicon.ico")

	// var database db.Database
	// database.Init(config)

	// r.Use(db.DB(&database))

	r.LoadHTMLGlob("templates/*")

	//routings
	r.GET("/", route.Main)
	r.GET("/ping", route.Ping)

	r.Run(":" + config.APP_PORT)
}
