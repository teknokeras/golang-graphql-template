package args

import (
	"github.com/graphql-go/graphql"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/types"
)

var Arguments = graphql.FieldConfigArgument{
    "user": &graphql.ArgumentConfig{
        Type: types.UserInputType,
    },
}