package deleterole

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

var Field = &graphql.Field{
			Type: graphql.String,
			Description: "Delete Role",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				if err := model.DeleteRole(params.Args["id"].(int)); err != nil {
					return nil, err
				}else{
					return "Role is Deleted", nil
				}
			},
		}