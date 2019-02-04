package schema

import "github.com/graphql-go/graphql"

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})
