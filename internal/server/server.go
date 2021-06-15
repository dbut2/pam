package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(address string) error {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	}

	root := r.Group("")
	{
		r.LoadHTMLGlob("templates/*")

		root.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "Pong!")
		})
	}

	return r.Run(address)
}
