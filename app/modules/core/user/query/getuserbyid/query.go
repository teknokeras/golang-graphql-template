package getuserbyid

import (
	"errors"

	"github.com/graphql-go/graphql"

	"app/graphqltypes"

	"app/modules/core/user/model"

)

var	Field = &graphql.Field{
	Type: graphqltypes.UserType, // the return type for this field
	Description: "Get User By ID",
	Args: graphql.FieldConfigArgument{
	    "id": &graphql.ArgumentConfig{
	        Type: graphql.Int,
	    },
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if (params.Args["id"] == nil){
			return nil, errors.New("ID argument is mandatory")
		}

		if r, err := model.GetUserById(params.Args["id"].(int)); err!=nil{
			return nil, errors.New("Such user doesn't exists")
		}else{
			return r, nil
		}
	},
} 