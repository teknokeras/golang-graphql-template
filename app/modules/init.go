package modules

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/orm"

	"github.com/teknokeras/golang-graphql-template/app/db"
	"github.com/teknokeras/golang-graphql-template/app/passwordutils"

	role "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
	user "github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
)

var ModelList = []interface{}{(*role.Role)(nil), (*user.User)(nil)}

func init() {

	db.Init()

	for _, model := range ModelList {
		err := db.Engine.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			panic(err)
		}
	}

	roleDoesntExists := true
	userDoesntExists := true

	r := role.Role{Name: os.Getenv("DEFAULT_ADMINISTRATOR_ROLE")}

	err := db.Engine.Select(&r)
	if err == nil {
		roleDoesntExists = false
	}

	var users []user.User
	err = db.Engine.Model(&users).Select()
	if err != nil {
		fmt.Println("we have a panic")
	}

	if len(users) > 0 {
		userDoesntExists = false
	}

	if (roleDoesntExists) && (userDoesntExists) {
		fmt.Println("Create fixture")

		err := db.Engine.Insert(&r)

		if err != nil {
			fmt.Println("Cannot create Default Role")
			panic(err)
		}

		encryptedPassword, errEnc := passwordutils.EncryptPassword(os.Getenv("DEFAULT_ADMIN_PASSWORD"))

		if errEnc != nil {
			panic("Cannot Encrypt Password")
		}

		u := user.User{
			Name:     os.Getenv("DEFAULT_ADMIN_FULL_NAME"),
			Email:    os.Getenv("DEFAULT_ADMIN_EMAIL"),
			Password: encryptedPassword,
			RoleId:   r.Id,
		}

		err = db.Engine.Insert(&u)

		if err != nil {
			fmt.Println("Cannot create Default User")
			panic(err)
		}
	}
}
