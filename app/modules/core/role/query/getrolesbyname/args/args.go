package args

import (
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/role/types"
)

var Arguments = graphql.FieldConfigArgument{
    "role": &graphql.ArgumentConfig{
        Type: types.RoleInputType,
    },
}