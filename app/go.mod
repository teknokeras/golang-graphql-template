module github.com/teknokeras/golang-graphql-template/app

replace github.com/teknokeras/golang-graphql-template/app/schema => ./schema

replace github.com/teknokeras/golang-graphql-template/app/graphqlutils => ./graphqlutils

replace github.com/teknokeras/golang-graphql-template/app/modules => ./modules

replace github.com/teknokeras/golang-graphql-template/app/modules/core/role => ./modules/core/role

replace github.com/teknokeras/golang-graphql-template/app/modules/core/user => ./modules/core/user

require (
	github.com/go-pg/pg v7.1.6+incompatible
	github.com/graph-gophers/graphql-go v0.0.0-20190108123631-d5b7dc6be53b
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/opentracing/opentracing-go v1.0.2 // indirect
	github.com/pkg/errors v0.8.1
	golang.org/x/crypto v0.0.0-20180910181607-0e37d006457b
	golang.org/x/net v0.0.0-20190119204137-ed066c81e75e // indirect
	golang.org/x/sys v0.0.0-20190116161447-11f53e031339 // indirect
	mellium.im/sasl v0.2.1 // indirect
)
