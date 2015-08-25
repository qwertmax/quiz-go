package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qwertmax/quiz-go/db"
)

func GetDB(c *gin.Context) *gorm.DB {
	database := c.MustGet("db").(*db.Database)
	return database.DB
}

func Main(c *gin.Context) {
	c.HTML(200, "leyout", gin.H{
		"test": "qqqqqq",
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"Message": "pong",
	})
}
