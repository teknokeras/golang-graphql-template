package createrole

import (
	//"errors"

	"github.com/graphql-go/graphql"

	"app/graphqltypes"
	"app/modules/core/role/model"
)

var Field = &graphql.Field{
			Type: graphqltypes.RoleType,
			Description: "Create new Role",
			Args: graphql.FieldConfigArgument{
				"role": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqltypes.RoleInputType),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				m := params.Args["role"].(map[string]interface{})
				if role, err := model.CreateRole(m["name"].(string)); err != nil {
					return model.Role{}, err
				}else{
					return role, nil
				}
			},
		}