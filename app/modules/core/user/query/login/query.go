package login

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/passwordutils"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
)

var Field = &graphql.Field{
	Type:        graphql.String,
	Description: "Login",
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		email := params.Args["email"].(string)
		password := params.Args["password"].(string)

		if model.VerifyPassword(email, password) {
			if token, err := passwordutils.GenerateJwtToken(email); err != nil {
				return nil, err
			} else {
				return token, nil
			}
		} else {
			return nil, errors.New("Email and/or Password are wrong")
		}
	},
}
