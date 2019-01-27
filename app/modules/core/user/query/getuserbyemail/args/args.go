package args

import (
	"github.com/graphql-go/graphql"
	"github.com/teknokeras/golang-graphql-template/app/modules/code/user/types"
)

var Arguments = graphql.FieldConfigArgument{
    "user": &graphql.ArgumentConfig{
        Type: types.UserInputType,
    },
}