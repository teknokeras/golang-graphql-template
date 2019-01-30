package getusersbyrole

import (
	"github.com/graphql-go/graphql"

	roleTypes "github.com/teknokeras/golang-graphql-template/app/modules/core/role/types"
)

var Arguments = graphql.FieldConfigArgument{
    "role": &graphql.ArgumentConfig{
        Type: roleTypes.RoleInputType,
    },
}