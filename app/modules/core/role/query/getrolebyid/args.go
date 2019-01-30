package getrolebyid

import "github.com/graphql-go/graphql"

var Arguments = graphql.FieldConfigArgument{
    "id": &graphql.ArgumentConfig{
        Type: graphql.Int,
    },
}