package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graphql-go/handler"

	auth "github.com/teknokeras/golang-graphql-template/app/auth"
	_ "github.com/teknokeras/golang-graphql-template/app/db"
	_ "github.com/teknokeras/golang-graphql-template/app/modules"
	appSchema "github.com/teknokeras/golang-graphql-template/app/schema"
)

func main() {

	r := chi.NewRouter()

	h := handler.New(&handler.Config{
		Schema:   &appSchema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r.Use(auth.IsAuthorized)

	r.Post("/graphql", h.ServeHTTP)

	fmt.Println("Now server is running on port 5000")
	fmt.Println("Login on      : curl -g 'http://localhost:5000/auth?mutation={Login(email:\"fsdfdf\",password:\"sdfdf\")}'")
	fmt.Println("Test with Get      : curl -g 'http://localhost:5000/graphql?query={Roles(first:1,after:2)}'")

	log.Fatal(http.ListenAndServe("0.0.0.0:5000", r))
}
