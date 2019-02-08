package db

import (
	"os"
	"strings"

	"github.com/go-pg/pg"
)

var Engine *pg.DB

func init() {
	Engine = newDB()
}

func newDB() *pg.DB {

	var dbHost strings.Builder

	dbHost.WriteString(os.Getenv("POSTGRES_HOST"))
	dbHost.WriteString(":")
	dbHost.WriteString(os.Getenv("POSTGRES_DB_PORT"))

	db := pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Addr:     dbHost.String(),
	})

	return db

}
