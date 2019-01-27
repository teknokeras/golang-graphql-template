package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/db"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getuserbyemail"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusers"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusersbyname"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/getusersbyrole"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getrolebyid"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getroles"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/query/getrolesbyname"
)

func RootQuery(database db.Database) graphql.NewObject {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//Role's queries
			"Roles": getrolebyid.GetField(database),
			"GetRoleById": getroles.GetField(database),
			"GetRoleByName": getrolesbyname.GetField(database),
			//User's queries
			"Users": getusers.GetField(database),
			"GetUsersByName": getusersbyname.GetField(database),
			"GetUserByEmail": getuserbyemail.GetField(database),
			"GetUsersByRole": getusersbyrole.GetField(database),
		},
	})	

	return rootQuery
}

