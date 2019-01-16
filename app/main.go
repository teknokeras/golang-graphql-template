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

	var u role.Role
	err = database.First(&u, 1).Error
	if err != nil {
		fmt.Println("Error getting role")
		fmt.Println(err)
	}
	fmt.Println(u.Name)

	//initDatabase(database)

	var schema = schema.BuildSchema()
	fmt.Println(schema)
}


func initDatabase(db *gorm.DB) {

	var role = role.Role{Name: "administrator"}

	err := db.Create(&role).Error

	if err != nil{
		return
	}

	err = db.Where("name=?","administrator").First(&role).Error

	if err != nil{
		return
	}

	fmt.Println(role.Name)
	fmt.Println(role.ID)

	var user = user.User{Name:"admin", Email:"admin@go.com", Password:"passwd", RoleID: role.ID}

	err = db.Create(&user).Error

	if err != nil{
		return
	}

	err = db.First(&user, 1).Error

	if err != nil{
		return
	}

	fmt.Println(user.Name)
	fmt.Println(user.Password)
	fmt.Println(user.Email)
	fmt.Println(user.ID)


}