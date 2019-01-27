package getuserbyemail

import (
	"errors"
	
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/user/types"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/user/args"
	"github.com/teknokeras/golang-graphql-template/app/modules/code/user/query/getroles/resolver"

)

var Field = &graphql.Field{
		Type: types.UserType, // the return type for this field
		Description: "Get User By Email",
		Args: args.Arguments,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			item, error := resolver.Resolve(params)
			return item, error
		},
	}