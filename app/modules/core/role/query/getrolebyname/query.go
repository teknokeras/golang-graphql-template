package getrolebyname

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/graphqltypes"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

var Field = &graphql.Field{
	Type:        graphqltypes.RoleType,
	Description: "Get Role By Name",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if params.Args["name"] == nil {
			return nil, errors.New("name argument is mandatory")
		}

		if r, err := model.GetRoleByName(params.Args["name"].(string)); err != nil {
			return nil, errors.New("Such role doesn't exists")
		} else {
			return r, nil
		}
	},
}
