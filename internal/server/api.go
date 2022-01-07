package server

import (
	"net/http"

	"github.com/dbut2/pam/internal/app"
	"github.com/dbut2/pam/pkg/models"
	"github.com/gin-gonic/gin"
)

func api(g *gin.RouterGroup, a app.App) {

	g.POST("/user", func(c *gin.Context) {

	})

	g.POST("/entry", func(c *gin.Context) {
		entry := &models.Entry{}
		err := c.BindJSON(&entry)
		if err != nil {
			panic(err.Error())
		}
		a := a.Authenticate(Authenticator(c))
		a.NewEntry()

	})

	g.GET("/entry/list", func(c *gin.Context) {
		a := a.Authenticate(Authenticator(c))
		entries := a.ListEntries()
		c.JSON(http.StatusOK, entries)
	})

	g.GET("/entry/:id", func(c *gin.Context) {
		a := a.Authenticate(Authenticator(c))
		id := c.Param("id")
		e := a.GetEntry(id)
		c.JSON(http.StatusOK, e)
	})

}
