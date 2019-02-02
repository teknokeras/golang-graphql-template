module github.com/teknokeras/golang-graphql-template/app

replace github.com/teknokeras/golang-graphql-template/app/schema => ./schema

replace github.com/teknokeras/golang-graphql-template/app/auth => ./auth

replace github.com/teknokeras/golang-graphql-template/app/modules => ./modules

replace github.com/teknokeras/golang-graphql-template/app/db => ./db

replace github.com/teknokeras/golang-graphql-template/app/graphqlutils => ./graphqlutils

replace github.com/teknokeras/golang-graphql-template/app/passwordutils => ./passwordutils

replace github.com/teknokeras/golang-graphql-template/app/typehub => ./typehub

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-pg/pg v7.1.7+incompatible
	github.com/graph-gophers/graphql-go v0.0.0-20190108123631-d5b7dc6be53b
	github.com/graphql-go/graphql v0.7.7
	github.com/graphql-go/handler v0.2.3
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/opentracing/opentracing-go v1.0.2 // indirect
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.3.0
	golang.org/x/crypto v0.0.0-20180910181607-0e37d006457b
	golang.org/x/net v0.0.0-20190125091013-d26f9f9a57f3 // indirect
	golang.org/x/sys v0.0.0-20190124100055-b90733256f2e // indirect
	mellium.im/sasl v0.2.1 // indirect
)
