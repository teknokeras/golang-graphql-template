package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/teknokeras/golang-graphql-template/app/db"
)

func GetSchema(database db.Database) graphql.SchemaConfig {
	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    RootQuery(database),
		Mutation: RootMutation(database),
	})

	return schema
}
