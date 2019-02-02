package updaterole

import (

	"github.com/graphql-go/graphql"

	"app/graphqltypes"
	"app/modules/core/role/model"
)

var Field = &graphql.Field{
			Type: graphqltypes.RoleType,
			Description: "Update Role",
			Args: graphql.FieldConfigArgument{
				"role": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqltypes.RoleInputType),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				args := params.Args["role"].(map[string]interface{})
				role := model.Role{Id: int64(args["id"].(int)), Name: args["name"].(string)}

				if role, err := model.UpdateRole(role); err != nil{
					return role, err
				}else{
					return role, nil
				}
			},
		}