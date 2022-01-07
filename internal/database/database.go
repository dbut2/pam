package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewDatabase(config Config) *DB {
	db := &DB{}

	dbConfig := mysql.Config{
		User:                 config.Username,
		Passwd:               config.Password,
		Net:                  "tcp",
		Addr:                 config.Hostname,
		DBName:               config.Database,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	conn, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	db.DB = conn

	return db
}
