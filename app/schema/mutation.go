package schema

import (
	"github.com/graphql-go/graphql"
	
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/mutation/createrole"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/mutation/updaterole"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/mutation/deleterole"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateRole": createrole.Field,
			"UpdateRole": updaterole.Field,
			"DeleteRole": deleterole.Field,
		},
   	})
