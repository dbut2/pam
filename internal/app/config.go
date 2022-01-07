package app

import (
	"github.com/dbut2/pam/internal/auth"
	"github.com/dbut2/pam/internal/database"
)

type Config struct {
	Database database.Config `yaml:"database"`
	Auth     auth.Config     `yaml:"auth"`
}
