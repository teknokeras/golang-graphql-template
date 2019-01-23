package modules

import (
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user"
)

var ModelList = []interface{}{(*role.Role)(nil), (*user.User)(nil)}
