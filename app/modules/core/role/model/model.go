package model

import (
	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	user "github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
)

type Role struct{
	Id     int64	
    Name   string	
    Users  []*user.User
}

type RoleList struct {
	PageInfo 	graphqlutils.PageInfo
	TotalCount 	int 
	Roles 		[]Role
}