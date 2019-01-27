package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/teknokeras/golang-graphql-template/app/db"
)

func RootMutation(database db.Database) graphql.NewObject {
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"CreateRole": ,
			"UpdateRole": ,
			"DeleteRole": ,
			"CreateUser": ,
			"UpdateUser": ,
			"DeleteUser": ,
   	})

	return rootMutation
}

