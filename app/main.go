package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"

	appSchema "github.com/teknokeras/golang-graphql-template/app/schema"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Print("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), appSchema.Schema)
		json.NewEncoder(w).Encode(result)
	})
}

/*
import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
                    "first": &graphql.ArgumentConfig{
                        Type: graphql.Int,
                    },
                },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if(p.Args["first"] == nil){
					log.Print("no first")
				}else{
					log.Print("first exists")
					log.Print(p.Args["first"].(int))
				}
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello(first: 2)
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}*/