package getrolebyid

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/graphqltypes"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

var Field = &graphql.Field{
	Type:        graphqltypes.RoleType,
	Description: "Get Role By ID",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if params.Args["id"] == nil {
			return nil, errors.New("'id' argument is mandatory")
		}

		if r, err := model.GetRoleById(params.Args["id"].(int)); err != nil {
			return nil, errors.New("Role Not Found")
		} else {
			return r, nil
		}

	},
}
