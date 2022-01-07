package app

import (
	"github.com/dbut2/pam/internal/auth"
	"github.com/dbut2/pam/internal/database"
	"github.com/dbut2/pam/pkg/models"
)

type App interface {
	Authenticator() *auth.Auth
	Authenticate(id string) AuthenticatedApp
	IsAuthenticated() bool

	NewUser() *models.User

	GetUser(id string) *models.User
	GetUserByGID(gid string) *models.User
}

type AuthenticatedApp interface {
	App
	NewEntry() *models.Entry
	NewMetadata() *models.Metadata

	ListEntries() []*models.Entry
	GetEntry(id string) *models.Entry

	UpdateEntry(entry *models.Entry)
}

type app struct {
	config Config
	db     *database.DB
	auth   *auth.Auth
}

func NewApp(config Config) App {
	a := &app{
		config: config,
	}

	a.db = database.NewDatabase(a.config.Database)
	a.auth = auth.NewAuth(a.config.Auth)

	return a
}

func (a *app) Authenticator() *auth.Auth {
	return a.auth
}

func (a *app) Authenticate(gid string) AuthenticatedApp {
	return &authenticatedApp{
		app:  a,
		user: a.GetUserByGID(gid),
	}
}

func (a app) IsAuthenticated() bool {
	return false
}

type authenticatedApp struct {
	*app
	user *models.User
}

func (a authenticatedApp) IsAuthenticated() bool {
	return true
}
