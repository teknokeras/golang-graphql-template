package user

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	RoleID   uint
	Name      string
	Email     string
	Password  string
}