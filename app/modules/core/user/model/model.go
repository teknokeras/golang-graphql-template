package model

import (
	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
)

type User struct{
	Id 		  int64
	Name      string
	Email     string
	Password  string
	RoleId	  int64
}

type UserList struct {
	PageInfo 	graphqlutils.PageInfo
	TotalCount 	int 
	Users 		[]User
}