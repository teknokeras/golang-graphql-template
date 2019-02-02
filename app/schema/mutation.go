package schema

import (
	"github.com/graphql-go/graphql"
	
	"app/modules/core/role/mutation/createrole"
	"app/modules/core/role/mutation/updaterole"
	"app/modules/core/role/mutation/deleterole"

	"app/modules/core/user/mutation/createuser"
	"app/modules/core/user/mutation/updateuser"
	"app/modules/core/user/mutation/deleteuser"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateRole": createrole.Field,
			"UpdateRole": updaterole.Field,
			"DeleteRole": deleterole.Field,
			"CreateUser": createuser.Field,
			"UpdateUser": updateuser.Field,
			"DeleteUser": deleteuser.Field,
		},
   	})
