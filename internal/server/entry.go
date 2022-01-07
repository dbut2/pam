package server

import (
	"net/http"

	"github.com/dbut2/pam/internal/app"
	"github.com/gin-gonic/gin"
)

func entry(g *gin.RouterGroup, a app.App) {
	g.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "entry", gin.H{})
	})

	g.POST("", func(c *gin.Context) {
		entry := c.PostForm("entry")
		a := a.Authenticate(Authenticator(c))
		e := a.NewEntry()

		e.Entry = entry

		a.UpdateEntry(e)
	})

	g.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")

		a := a.Authenticate(Authenticator(c))

		a.GetEntry(id)

		c.HTML(http.StatusOK, "entry", gin.H{})
	})

	g.GET("/list", func(c *gin.Context) {
		a := a.Authenticate(Authenticator(c))
		entries := a.ListEntries()

		c.HTML(http.StatusOK, "entries", gin.H{
			"entries": entries,
		})
	})
}
