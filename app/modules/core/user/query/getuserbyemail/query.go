package getuserbyemail

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/graphqltypes"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"

)

var	Field = &graphql.Field{
	Type: graphqltypes.UserType, // the return type for this field
	Description: "Get User By Email",
	Args: graphql.FieldConfigArgument{
	    "email": &graphql.ArgumentConfig{
	        Type: graphql.String,
	    },
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if (params.Args["email"] == nil){
			return nil, errors.New("email argument is mandatory")
		}

		if r, err := model.GetUserByEmail(params.Args["email"].(string)); err!=nil{
			return nil, errors.New("Such user doesn't exists")
		}else{
			return r, nil
		}
	},
} 