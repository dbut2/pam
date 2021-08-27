package server

import (
	"html/template"
	"net/http"

	"github.com/dbut2/pam/templates"
	"github.com/gin-gonic/gin"
)

func Run(address string) error {

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/user/{user}", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})

		api.GET("/:user", func(c *gin.Context) {
			code := c.Param("code")
			s := a.Lengthen(code)
			c.Redirect(http.StatusTemporaryRedirect, s.Url)
		})
	}

	root := r.Group("")
	{
		r.GET("index", func(c *gin.Context) {
			t, err := template.New("index").Parse(templates.Index)
			_, _ = t, err
		})

		root.GET("/cal", func(c *gin.Context) {

		})

		root.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "Pong!")
		})
	}

	return r.Run(address)
}
