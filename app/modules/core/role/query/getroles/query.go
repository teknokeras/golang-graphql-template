package getroles

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/graphqltypes"
	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

var Field = &graphql.Field{
	Type:        graphqltypes.RoleListType,
	Description: "Get All Roles paginated",
	Args: graphql.FieldConfigArgument{
		"first": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"after": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		first := 10
		after := 0

		if params.Args["first"] != nil {
			first, _ = params.Args["first"].(int)
		}

		if params.Args["after"] != nil {
			after, _ = params.Args["after"].(int)
		}

		if roles, err := model.GetRolesPaginated(first, after); err != nil {
			return nil, errors.New("No Roles found")
		} else {
			isPrevExists, isNextExists := model.GetPrevNextExistence(after, roles)

			pageInfo := graphqlutils.PageInfo{HasNextPage: isNextExists, HasPreviousPage: isPrevExists}

			roleList := model.RoleList{TotalCount: len(roles), PageInfo: pageInfo, Roles: roles}

			return roleList, nil

		}

	},
}
