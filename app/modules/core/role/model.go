package role

import (
	"github.com/jinzhu/gorm"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user"
)

type Role struct{
	gorm.Model
	Name string
	Users []user.User `gorm:"foreignkey:RoleID"`
}