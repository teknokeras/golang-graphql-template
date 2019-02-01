package updateuser

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/graphqltypes"
	roleModel "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
)

var Field = &graphql.Field{
			Type: graphqltypes.UserType,
			Description: "Update User",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqltypes.UserInputType),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				m := params.Args["user"].(map[string]interface{})

				userId := m["id"].(int)

				if user, err := model.GetUserById(userId); err != nil {
					return nil, errors.New("Cannot find user")
				}else{

					roleId := m["roleId"].(int)

					if role, err := roleModel.GetRoleById(roleId); err != nil {
						return nil, errors.New("Wrong Role ID")
					}else{
						user.Name = m["name"].(string)
						user.Email = m["email"].(string)
						user.RoleId = role.Id

						if user, err := model.UpdateUser(user); err != nil {
							return nil, err
						}else{
							return user, nil
						}
					}
				} 
			},
		}