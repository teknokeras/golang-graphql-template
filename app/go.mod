module github.com/teknokeras/golang-graphql-template/app

replace github.com/teknokeras/golang-graphql-template/app/schema => ./schema

require (
	github.com/go-pg/pg v7.1.7+incompatible
	github.com/graphql-go/graphql v0.7.7
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	mellium.im/sasl v0.2.1 // indirect
)
