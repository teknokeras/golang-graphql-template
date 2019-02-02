package schema

import (
	"github.com/graphql-go/graphql"

	"app/modules/core/role/query/getroles"
	"app/modules/core/role/query/getrolebyid"
	"app/modules/core/role/query/getrolebyname"

	"app/modules/core/user/query/getuserbyemail"
	"app/modules/core/user/query/getuserbyid"
	"app/modules/core/user/query/getusers"
	"app/modules/core/user/query/getusersbyname"
	"app/modules/core/user/query/getusersbyrole"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"GetRoles": getroles.Field,
		"GetRoleById": getrolebyid.Field,
		"GetRoleByName": getrolebyname.Field,
		"GetUserByEmail": getuserbyemail.Field,
		"GetUserById": getuserbyid.Field,
		"GetUsers": getusers.Field,
		"GetUsersByName": getusersbyname.Field,
		"GetUsersByRole": getusersbyrole.Field,
	},
})	
