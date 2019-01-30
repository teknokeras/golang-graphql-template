package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getroles"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getrolebyid"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getrolebyname"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"Roles": getroles.Field,
		"GetRoleById": getrolebyid.Field,
		"GetRoleByName": getrolebyname.Field,
	},
})	
