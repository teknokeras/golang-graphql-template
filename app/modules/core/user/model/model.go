package model

import (
	"errors"
	
	"app/passwordutils"
	"app/graphqlutils"
	"app/db"
)

type User struct{
	Id 		  int64
	Name      string
	Email     string
	Password  string
	RoleId	  int64
}

type UserList struct {
	PageInfo 	graphqlutils.PageInfo
	TotalCount 	int 
	Users 		[]User
}

func GetUserById(id int) (User, error){
	u := User{Id: int64(id)}

	err := db.Engine.Select(&u)
	if err != nil {
		return User{}, errors.New("User Not Found")
    }
    return u, nil
}

func GetUserByEmail(email string) (User, error){
	r := User{}

	err := db.Engine.Model(&r).Where("email=?", email).Select()
	if err != nil {
		return User{}, errors.New("User Not Found")
    }
    return r, nil
}

func GetUsersPaginated(first int, after int) ([]User, error){
	var users []User
	err := db.Engine.Model(&users).Where("id>?", after).Limit(first).Order("id ASC").Select()

	if err != nil{
		return users, errors.New("Users Not Found")
	}

	if (len(users) == 0){
		return users, errors.New("No Users match the criteria are found")
	}

	return users, nil
}

func GetUsersByNamePaginated(first int, after int, name string) ([]User, error){
	var users []User
	err := db.Engine.Model(&users).Where("name LIKE ?", "%"+name+"%").Where("id>?", after).Limit(first).Order("id ASC").Select()

	if err != nil{
		return users, errors.New("Users Not Found")
	}

	if (len(users) == 0){
		return users, errors.New("No Users match the criteria are found")
	}

	return users, nil
}

func GetUsersByRolePaginated(first int, after int, roleId int) ([]User, error){
	var users []User
	err := db.Engine.Model(&users).Where("roleId=?", roleId).Where("id>?", after).Limit(first).Order("id ASC").Select()

	if err != nil{
		return users, errors.New("Users Not Found")
	}

	if (len(users) == 0){
		return users, errors.New("No Users match the criteria are found")
	}

	return users, nil
}

func GetPrevNextExistence(users []User) (bool, bool){
	user := users[0]
	prevCount, _ := db.Engine.Model((*User)(nil)).Where("id < ?", user.Id).Order("id ASC").Count()
	
	user = users[len(users) - 1]
	nextCount, _ := db.Engine.Model((*User)(nil)).Where("id > ?", user.Id).Order("id ASC").Count()

	return prevCount > 0, nextCount > 0
}

func CreateUser(name string, email string, password string, roleId int) (User, error){

	if u, err := GetUserByEmail(email); err == nil{
		//user already exists 
		return User{}, errors.New("User already exists")
	}else{
		if hashedPassword, errEnc := passwordutils.EncryptPassword(password); errEnc != nil {
			return User{}, errors.New("Cannot hash password")
		}else{
			u = User{Name: name, Email: email, Password: hashedPassword, RoleId: int64(roleId)}

			if err = db.Engine.Insert(&u); err != nil{
				return User{}, err
			}
			return u, nil
		}
	}
}

func UpdateUser(user User) (User, error){
	r := User{Id: user.Id}

	if err := db.Engine.Select(&r); err != nil{
		return User{}, err
	}else{
		r.Name = user.Name
		r.Email = user.Email
		r.RoleId = user.RoleId

		if err = db.Engine.Update(&r); err != nil{
			return User{}, err
		}else{
			return r, nil
		}
	}
}

func DeleteUser(id int) (error) {
	r := User{Id: int64(id)}

	if err := db.Engine.Select(&r); err != nil{
		return errors.New("User Does not Exists")
	}else{
		if err = db.Engine.Delete(&r); err != nil{
			return errors.New("User Cannot be deleted")
		}else{
			return nil
		}
	}
}

func VerifyPassword(email string, password string) bool {
	
	if u, err := GetUserByEmail(email); err != nil{
		//user already exists 
		return false
	}else{
		//user exists
		if match, err := passwordutils.ComparePasswordAndHash(password, u.Password); err != nil {
			return false
		} else {
			if match {
				return true
			}else{
				return false
			}
		}
	}
}