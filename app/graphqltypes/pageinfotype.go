package graphqltypes

import "github.com/graphql-go/graphql"

var PageInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PageInfo",
	Description: "Information about pagination in a list.",
	Fields: graphql.Fields{
		"hasNextPage": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Boolean),
			Description: "When paginating forwards, are there more items?",
		},
		"hasPreviousPage": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Boolean),
			Description: "When paginating backwards, are there more items?",
		},
	},
})
