package getusers

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	"github.com/teknokeras/golang-graphql-template/app/graphqltypes"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
)

var Field = &graphql.Field{
	Type: graphqltypes.UserListType,
	Description: "Get All Users paginated",
	Args: graphql.FieldConfigArgument{
	    "first": &graphql.ArgumentConfig{
	        Type: graphql.Int,
	    },
	    "after": &graphql.ArgumentConfig{
	        Type: graphql.Int,
	    },
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		first := 10
		after := 0

		if (params.Args["first"] != nil){
			first, _ = params.Args["first"].(int)
		}

		if (params.Args["after"] != nil){
			after, _ = params.Args["after"].(int)
		}

		if users, err := model.GetUsersPaginated(first, after); err != nil{
			return nil, errors.New("No Users found")
		}else{
			isPrevExists, isNextExists := model.GetPrevNextExistence(users)

			pageInfo := graphqlutils.PageInfo{HasNextPage: isNextExists, HasPreviousPage: isPrevExists}

			userList := model.UserList{TotalCount: len(users), PageInfo: pageInfo, Users: users}

			return userList, nil
			
		}

	},
} 