package args

import (
	"github.com/graphql-go/graphql"
	"app/modules/core/user/types"
)

var Arguments = graphql.FieldConfigArgument{
    "user": &graphql.ArgumentConfig{
        Type: types.UserInputType,
    },
}