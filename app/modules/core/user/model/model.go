package model

import role "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"

type User struct{
	Id 		  int64
	Name      string
	Email     string
	Password  string
	Role      *role.Role
}