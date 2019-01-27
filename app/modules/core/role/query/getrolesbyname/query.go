package getrolebyname

import (
	"errors"
	
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/role/types"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/role/args"
	"github.com/teknokeras/golang-graphql-template/app/modules/code/role/query/getrolesbyname/resolver"

)

var Field = &graphql.Field{
		Type: types.RoleList, // the return type for this field
		Description: "Get All Roles By Name paginated",
		Args: args.Arguments,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			item, error := resolver.Resolve(params)
			return item, error
		},
	}