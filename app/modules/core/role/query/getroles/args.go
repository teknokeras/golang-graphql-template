package getroles

import "github.com/graphql-go/graphql"

var Arguments = graphql.FieldConfigArgument{
    "first": &graphql.ArgumentConfig{
        Type: graphql.Int,
    },
    "after": &graphql.ArgumentConfig{
        Type: graphql.Int,
    },
}

