package server

import (
	"net/http"
	"time"

	"github.com/dbut2/pam/internal/app"
	"github.com/dbut2/pam/internal/server/pages"
	"github.com/dbut2/pam/internal/server/templates"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config Config
}

func NewServer(config Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Serve(a app.App) error {
	if s.config.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	t, err := templates.GetTemplate()
	if err != nil {
		return err
	}
	r.SetHTMLTemplate(t)

	api(r.Group("/api"), a)
	entry(r.Group("/entry"), a)

	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(pages.Site))
	})

	r.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, a.Authenticator().GetLoginURL())
	})

	r.GET("/auth", func(c *gin.Context) {
		code := c.Query("code")
		ts := a.Authenticator().Exchange(c, code)

		gid := a.Authenticator().GetGID(c, ts)

		cookie := &http.Cookie{
			Name:     "GID",
			Value:    gid,
			MaxAge:   int((time.Hour * 24 * 365).Seconds()),
			Secure:   true,
			HttpOnly: true,
		}

		http.SetCookie(c.Writer, cookie)

		c.Redirect(http.StatusTemporaryRedirect, "")
	})

	r.GET("/logout", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "GID",
			MaxAge:   -1,
			Secure:   true,
			HttpOnly: true,
		}

		http.SetCookie(c.Writer, cookie)

		c.Redirect(http.StatusTemporaryRedirect, "")
	})

	return r.Run(s.config.Address)
}

func Serve(config Config, a app.App) error {
	return NewServer(config).Serve(a)
}

func Authenticator(c *gin.Context) string {
	gid, err := c.Cookie("GID")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return gid
}
