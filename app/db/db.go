package db

import (
	"os"
	"strings"
	"github.com/jinzhu/gorm"
	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func buildPath() string{
	var path strings.Builder

	path.WriteString("host=")
	path.WriteString(os.Getenv("HOST"))
	path.WriteString(" user="+os.Getenv("USER"))
	path.WriteString(" dbname="+os.Getenv("DBNAME"))
	path.WriteString(" sslmode=disable password=")
	path.WriteString(os.Getenv("PASSWORD"))

	return path.String()
}

func NewDatabase() (*gorm.DB, error) {
	// connect to the example db, create if it doesn't exist
	db, err := gorm.Open("postgres", buildPath())
  	
	if err != nil {
		return nil, err
	}

	return db, err
}