package main

import (
    "net/http"

    log "github.com/sirupsen/logrus"
    graphql "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"

    schemaBuilder "github.com/teknokeras/golang-graphql-template/app/schema"
    "github.com/teknokeras/golang-graphql-template/app/db"
)

func main() {
    schemaString := schemaBuilder.Build()

    database := db.NewDB()

    if (database == nil){
        panic("Cannot create database")
    }

    schema := graphql.MustParseSchema(s, &schema.Resolver{Db: database}, graphql.UseStringDescriptions())

    mux := http.NewServeMux()

    mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write(page)
    }))

    mux.Handle("/query", &relay.Handler{Schema: schema})

    log.WithFields(log.Fields{"time": time.Now()}).Info("starting server")
    log.Fatal(http.ListenAndServe("0.0.0.0:5000", logged(mux)))
}

// logging middleware
func logged(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now().UTC()

        next.ServeHTTP(w, r)

        log.WithFields(log.Fields{
            "path":    r.RequestURI,
            "IP":      r.RemoteAddr,
            "elapsed": time.Now().UTC().Sub(start),
        }).Info()
    })
}
