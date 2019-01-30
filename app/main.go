package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"

	appSchema "github.com/teknokeras/golang-graphql-template/app/schema"
	"github.com/teknokeras/golang-graphql-template/app/modules"
)


func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {

	modules.InitTablesAndFixtures()

    schema := appSchema.Schema

    http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 5000")
	fmt.Println("Test with Get      : curl -g 'http://localhost:5000/graphql?query={Roles(first:1,after:2)}'")
	http.ListenAndServe("0.0.0.0:5000", nil)

}



