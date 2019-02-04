package deleteuser

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
)

var Field = &graphql.Field{
	Type:        graphql.String,
	Description: "Delete User",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		if err := model.DeleteUser(params.Args["id"].(int)); err != nil {
			return nil, err
		} else {
			return "User is Deleted", nil
		}
	},
}
