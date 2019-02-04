package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getrolebyid"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getrolebyname"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getroles"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getuserbyemail"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getuserbyid"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusers"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusersbyname"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusersbyrole"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/login"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"GetRoles":       getroles.Field,
		"GetRoleById":    getrolebyid.Field,
		"GetRoleByName":  getrolebyname.Field,
		"GetUserByEmail": getuserbyemail.Field,
		"GetUserById":    getuserbyid.Field,
		"GetUsers":       getusers.Field,
		"GetUsersByName": getusersbyname.Field,
		"GetUsersByRole": getusersbyrole.Field,
		"Login":          login.Field,
	},
})
