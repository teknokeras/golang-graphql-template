package createuser

import (
	"errors"

	"github.com/graphql-go/graphql"

	"app/graphqltypes"
	roleModel "app/modules/core/role/model"
	"app/modules/core/user/model"
)

var Field = &graphql.Field{
			Type: graphqltypes.UserType,
			Description: "Create new User",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqltypes.UserInputType),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				m := params.Args["user"].(map[string]interface{})

				roleId := m["roleId"].(int)

				if role, err := roleModel.GetRoleById(roleId); err != nil {
					return nil, errors.New("Role not found")
				}else{
					if user, err := model.CreateUser(m["name"].(string), m["email"].(string), m["password"].(string), int(role.Id)); err != nil {
						return model.User{}, err
					}else{
						return user, nil
					}
				}
			},
		}