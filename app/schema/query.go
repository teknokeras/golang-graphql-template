package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getuserbyemail"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusers"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusersbyname"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusersbyrole"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getrolebyid"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getroles"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getrolesbyname"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		//Role's queries
		"Roles": getrolebyid.Field,
		"GetRoleById": getroles.Field,
		"GetRoleByName": getrolesbyname.Field,
		//User's queries
		"Users": getusers.Field,
		"GetUsersByName": getusersbyname.Field,
		"GetUserByEmail": getuserbyemail.Field,
		"GetUsersByRole": getusersbyrole.Field,
	},
})