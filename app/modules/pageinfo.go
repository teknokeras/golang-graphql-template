package modules

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
		"startCursor": &graphql.Field{
			Type:        graphql.String,
			Description: "When paginating backwards, the cursor to continue.",
		},
		"endCursor": &graphql.Field{
			Type:        graphql.String,
			Description: "When paginating forwards, the cursor to continue.",
		},
	},
})