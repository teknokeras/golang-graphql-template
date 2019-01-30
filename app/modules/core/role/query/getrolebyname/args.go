package getrolebyname

import "github.com/graphql-go/graphql"

var Arguments = graphql.FieldConfigArgument{
    "name": &graphql.ArgumentConfig{
        Type: graphql.String,
    },
}