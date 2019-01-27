package getusersbyrole

import (
	"errors"
	
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/user/types"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/user/query/getrolesbyname/args"
	"github.com/teknokeras/golang-graphql-template/app/modules/code/user/query/getrolesbyname/resolver"

)

var Field = &graphql.Field{
		Type: types.UserList, // the return type for this field
		Description: "Get All Users By Role",
		Args: args.Arguments,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			item, error := resolver.Resolve(params)
			return item, error
		},
	}