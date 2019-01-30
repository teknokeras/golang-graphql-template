package getrolebyid

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/db"

	"github.com/teknokeras/golang-graphql-template/app/typehub"
)


func GetField(database *db.Database) *graphql.Field{
	field := &graphql.Field{
		Type: typehub.BuildRoleType(database), // the return type for this field
		Description: "Get Role By ID",
		Args: Arguments,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			item, error := Resolve(params, database)
			return item, error
		},
	} 

	return field
}
