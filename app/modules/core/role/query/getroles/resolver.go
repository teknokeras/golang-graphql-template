package getroles

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/db"

	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

func Resolve(params graphql.ResolveParams) (interface{}, error){
	first := 10
	after := 0

	if (params.Args["first"] != nil){
		first, _ = params.Args["first"].(int)
	}

	if (params.Args["after"] != nil){
		after, _ = params.Args["after"].(int)
	}

	var roles []model.Role
	err := db.Engine.Model(&roles).Where("id > ?", after).Order("id ASC").Limit(first).Select()

	if err != nil{
		return nil, errors.New("Role Not Found")
	}

	if (len(roles) == 0){
		return nil, errors.New("No Roles match the criteria are found")
	}

	prevCount, _ := db.Engine.Model((*graphqlutils.PageInfo)(nil)).Where("id < ?", after).Order("id ASC").Count()
	
	nextCount, _ := db.Engine.Model((*graphqlutils.PageInfo)(nil)).Where("id > ?", roles[len(roles) - 1].Id).Order("id ASC").Count()

	pageInfo := graphqlutils.PageInfo{HasNextPage: nextCount > 0, HasPreviousPage: prevCount > 0}

	roleList := model.RoleList{TotalCount: len(roles), PageInfo: pageInfo, Roles: roles}

	return roleList, nil
}