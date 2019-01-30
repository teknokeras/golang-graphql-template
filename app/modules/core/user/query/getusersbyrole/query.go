package getusersbyrole

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/typehub"
	"github.com/teknokeras/golang-graphql-template/app/db"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/args"

)

func GetField(database db.Database) *graphql.Field{
	field := &graphql.Field{
		Type: typehub.GetUserListType(database), // the return type for this field
		Description: "Get All Users By Role",
		Args: args.Arguments,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			item, error := Resolve(params, database)
			return item, error
		},
	} 

	return field
}
