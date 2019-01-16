package db

import (
	"os"
	"fmt"
	"strings"
	"github.com/jinzhu/gorm"
	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user"

)

func buildPath() string{
	var path strings.Builder

	path.WriteString("host=")
	path.WriteString(os.Getenv("POSTGRES_HOST"))
	path.WriteString(" user="+os.Getenv("POSTGRES_USER"))
	path.WriteString(" dbname="+os.Getenv("POSTGRES_DB"))
	path.WriteString(" sslmode=disable password=")
	path.WriteString(os.Getenv("POSTGRES_PASSWORD"))

	return path.String()
}

func NewDatabase() (*gorm.DB, error) {
	var path = buildPath()

	fmt.Println(path)

	// connect to the example db, create if it doesn't exist
	db, err := gorm.Open("postgres", buildPath())
  	
	if err != nil {
		return nil, err
	}

	db.DropTableIfExists(&user.User{}, &role.Role{})
	db.AutoMigrate(&user.User{}, &role.Role{})

	return db, err
}