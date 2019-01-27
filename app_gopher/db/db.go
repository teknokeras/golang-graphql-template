package db

import (
	"os"
	"strings"

    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"

	"github.com/teknokeras/golang-graphql-template/app/modules"
)

// DB is the DB that will performs all operation
type Database struct {
	Engine *pg.DB
}

func NewDB() *Database{

	var dbHost strings.Builder

	dbHost.WriteString(os.Getenv("POSTGRES_HOST"))
	dbHost.WriteString(":")
	dbHost.WriteString(os.Getenv("POSTGRES_DB_PORT"))
	
	db := pg.Connect(&pg.Options{
        User: os.Getenv("POSTGRES_USER"),
        Password: os.Getenv("POSTGRES_PASSWORD"),
        Database: os.Getenv("POSTGRES_DB"),
        Addr: dbHost.String(),
    })

    for _, model := range modules.ModelList {
        err := db.CreateTable(model, &orm.CreateTableOptions{
            IfNotExists: true,
        })
        if err != nil {
            return nil
        }
    }

    modules.InitFixtures(db)

	return &Database{db}

}