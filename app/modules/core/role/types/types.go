package graphqltype

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules"
)

var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RoleList = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role List",
	Fields: graphql.Fields{
		"roles": &graphql.Field{
			Type:        graphql.NewList(RoleType),
			Description: "Roles of the list.",
		},
		"pageInfo": &graphql.Field{
			Type:        graphql.NewNonNull(modules.PageInfoType),
			Description: "Information to aid in pagination.",
		},
		"totalCount": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Count of all list roles.",
		},
	},
})

var RoleInputType = graphql.NewInputObject(
    graphql.InputObjectConfig{
        Name: "Role Input Type",
        Fields: graphql.InputObjectConfigFieldMap{
            "name": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
        },
    },
)