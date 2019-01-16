package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/teknokeras/golang-graphql-template/app/schema"
	"github.com/teknokeras/golang-graphql-template/app/db"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user"
)

func main() {
	var database, err = db.NewDatabase()

	if err != nil {
		fmt.Println("Database cannot be initiated")
		fmt.Println(err)
	}

	initDatabase(database)

	var schema = schema.BuildSchema()
	fmt.Println(schema)
}

func initDatabase(db *gorm.DB) {
	*db.DropTableIfExists(&user.User{}, &role.Role{})
	*db.AutoMigrate(&user.User{}, &role.Role{})

	var role = role.Role{Name: "administrator"}

	err := *db.Create(&role).Error

	if err != nil{
		return
	}

	err = *db.Where("name = ?", "administrator").First(&role).Error
	if err != nil {
		return nil
	}

	var admin = user.User{Name: "admin", email:"admin@go.com", password:"testing", RoleID: &role.id}

	err = *db.Create(&admin).Error

	if err != nil{
		return
	}

}