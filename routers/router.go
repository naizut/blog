package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	router := gin.Default();

	router.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Bonjour, Je suis Marchelli %s", name)
	})
	
	router.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Bonjour, Je suis Marchelli12")
	})
	
	router.Run(":8080")
}