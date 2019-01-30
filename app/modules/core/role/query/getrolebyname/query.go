package getrolebyname

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/graphqltypes"
)


var Field = &graphql.Field{
	Type: graphqltypes.RoleType,
	Description: "Get Role By Name",
	Args: Arguments,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		item, error := Resolve(params)
		return item, error
	},
} 

