package modules

import (
	"os"
	"fmt"
    
    "github.com/go-pg/pg"

	"github.com/teknokeras/golang-graphql-template/app/passwordutils"

	role "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
	user "github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
)

var ModelList = []interface{}{(*role.Role)(nil), (*user.User)(nil)}

func InitFixtures(db *pg.DB){

	roleDoesntExists := true
	userDoesntExists := true

	r := role.Role{Name: os.Getenv("DEFAULT_ADMINISTRATOR_ROLE")}


	err := db.Select(&r)
	if err == nil {
		roleDoesntExists = false
    }

    fmt.Println(r)

    var users []user.User
    err = db.Model(&users).Select()
    if err != nil {
        fmt.Println("we have a panic")
    }

    if(len(users) > 0){
    	userDoesntExists = false
    }

    fmt.Println(users)

    if(roleDoesntExists) && (userDoesntExists){
    	fmt.Println("Create fixture")

    	err := db.Insert(&r)

		if err != nil{
			fmt.Println("Cannot create Default Role")
			panic(err)
			return
		}

		encryptedPassword, errEnc := passwordutils.EncryptPassword(os.Getenv("DEFAULT_ADMIN_PASSWORD"))

		if (errEnc != nil){
			panic("Cannot Encrypt Password")
		}

		u := user.User{
			Name: os.Getenv("DEFAULT_ADMIN_FULL_NAME"), 
			Email: os.Getenv("DEFAULT_ADMIN_EMAIL"), 
			Password: encryptedPassword, 
			Role: &r,
		}

		err = db.Insert(&u)

		if err != nil{
			fmt.Println("Cannot create Default User")
			panic(err)
		}
    }
}
