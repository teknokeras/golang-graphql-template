package user

import "github.com/teknokeras/golang-graphql-template/app/modules/core/role"

type User struct{
	Id 		  int64
	Name      string
	Email     string
	Password  string
	Role      *role.Role
}
/*
import "github.com/jinzhu/gorm"


type User struct{
	gorm.Model
	Name      string
	Email     string
	Password  string
	RoleID    int
}
*/