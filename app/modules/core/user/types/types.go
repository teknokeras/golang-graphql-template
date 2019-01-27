package types

package graphqltype

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserList = graphql.NewObject(graphql.ObjectConfig{
	Name: "User List",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(UserType),
			Description: "Users of the list.",
		},
		"pageInfo": &graphql.Field{
			Type:        graphql.NewNonNull(modules.PageInfoType),
			Description: "Information to aid in pagination.",
		},
		"totalCount": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Count of all list users.",
		},
	},
})

var UserInputType = graphql.NewInputObject(
    graphql.InputObjectConfig{
        Name: "User Input Type",
        Fields: graphql.InputObjectConfigFieldMap{
            "name": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
            "email": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
            "password": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
        },
    },
)