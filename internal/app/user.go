package app

import (
	"github.com/dbut2/pam/pkg/models"
	"github.com/google/uuid"
)

func (a *app) NewUser() *models.User {
	u := &models.User{
		ID: uuid.New().String(),
	}

	a.StoreUser(u)

	return u
}

func (a *app) StoreUser(user *models.User) {
	result, err := a.db.Exec("INSERT INTO user (id) VALUES (?)", user.ID)
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
}

func (a *app) GetUser(id string) *models.User {
	rows, err := a.db.Query("SELECT id, gid FROM user WHERE id = ? LIMIT 1", id)
	if err != nil {
		panic(err.Error())
	}

	user := &models.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.GID)
		if err != nil {
			panic(err.Error())
		}
	}

	return user
}

func (a *app) GetUserByGID(gid string) *models.User {
	rows, err := a.db.Query("SELECT id, gid FROM user WHERE gid = ? LIMIT 1", gid)
	if err != nil {
		panic(err.Error())
	}

	user := &models.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.GID)
		if err != nil {
			panic(err.Error())
		}
	}

	return user
}
