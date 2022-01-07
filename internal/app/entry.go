package app

import (
	_ "embed"
	"time"

	"github.com/dbut2/pam/pkg/models"
	"github.com/google/uuid"
)

func (a *authenticatedApp) NewEntry() *models.Entry {
	e := &models.Entry{
		ID:    uuid.New().String(),
		Date:  time.Now().Local(),
		Entry: "",
	}

	a.StoreEntry(e)

	return e
}

func (a *authenticatedApp) StoreEntry(entry *models.Entry) {
	result, err := a.db.Exec("INSERT INTO entry (id) VALUES (?)", entry.ID)
	if err != nil {
		panic(err.Error())
	}

	rows, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	if rows != 1 {
		panic("rows inserted not 1")
	}

	a.UpdateEntry(entry)
}

//go:embed queries/latestentry.sql
var latestEntrySQL string

func (a *authenticatedApp) ListEntries() []*models.Entry {
	rows, err := a.db.Query(latestEntrySQL, a.user.ID)
	if err != nil {
		panic(err.Error())
	}

	entries := []*models.Entry{}
	for rows.Next() {
		entry := &models.Entry{}
		err = rows.Scan(&entry.ID, &entry.Date, &entry.Entry)
		if err != nil {
			panic(err.Error())
		}
		entries = append(entries, entry)
	}

	return entries
}

func (a *authenticatedApp) GetEntry(id string) *models.Entry {
	rows, err := a.db.Query("SELECT e.id, eh.date, eh.entry FROM entry AS e LEFT JOIN entry_history eh on eh.id = e.id LEFT JOIN metadata m on m.id = eh.metadata WHERE e.id = ? ORDER BY m.timestamp DESC LIMIT 1", id)
	if err != nil {
		panic(err.Error())
	}

	entry := &models.Entry{}
	for rows.Next() {
		err = rows.Scan(&entry.ID, &entry.Date, &entry.Entry)
		if err != nil {
			panic(err.Error())
		}
	}

	return entry
}

func (a *authenticatedApp) UpdateEntry(entry *models.Entry) {
	metadata := a.NewMetadata()

	_, offset := time.Now().Zone()
	date := time.Now().UTC().Add(time.Duration(offset) * time.Second)

	result, err := a.db.Exec("INSERT INTO entry_history (id, metadata, date, entry) VALUES (?, ?, ?, ?)", entry.ID, metadata.ID, date, entry.Entry)
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
