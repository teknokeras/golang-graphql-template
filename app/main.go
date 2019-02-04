package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"

	auth "github.com/teknokeras/golang-graphql-template/app/auth"
	"github.com/teknokeras/golang-graphql-template/app/modules"
	appSchema "github.com/teknokeras/golang-graphql-template/app/schema"
)

func main() {
	modules.InitTablesAndFixtures()

	h := handler.New(&handler.Config{
		Schema:   &appSchema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	authHandler := handler.New(&handler.Config{
		Schema:   &authSchema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/auth", authHandler)
	http.Handle("/graphql", auth.IsAuthorized(h))

	fmt.Println("Now server is running on port 5000")
	fmt.Println("Login on      : curl -g 'http://localhost:5000/auth?mutation={Login(email:\"fsdfdf\",password:\"sdfdf\")}'")
	fmt.Println("Test with Get      : curl -g 'http://localhost:5000/graphql?query={Roles(first:1,after:2)}'")

	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
