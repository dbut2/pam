package app

import (
	"time"

	"github.com/dbut2/pam/pkg/models"
	"github.com/google/uuid"
)

func (a *authenticatedApp) NewMetadata() *models.Metadata {
	md := &models.Metadata{
		ID:        uuid.New().String(),
		Author:    a.user.ID,
		Timestamp: time.Now(),
	}

	a.StoreMetadata(md)

	return md
}

func (a *app) StoreMetadata(metadata *models.Metadata) {
	result, err := a.db.Exec("INSERT INTO metadata (id, author, timestamp) VALUES (?, ?, ?)", metadata.ID, metadata.Author, metadata.Timestamp)
	if err != nil {
		panic(err.Error())
	}

	rows, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	if rows != 1 {
		panic("inserted rows not 1")
	}
}
