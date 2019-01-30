package args

import (
	"github.com/graphql-go/graphql"

	types "github.com/teknokeras/golang-graphql-template/app/modules/core/role/types"
)

var Arguments = graphql.FieldConfigArgument{
    "role": &graphql.ArgumentConfig{
        Type: types.RoleInputType,
    },
}