package schema

import (
	"github.com/graphql-go/graphql"
	
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/query/login"
)

var AuthQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Login": login.Field,
		},
   	})
