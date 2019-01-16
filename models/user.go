package models

import (
	"crypto/md5"
	"fmt"
)

type IUserOperation interface {
	CellphoneLogin() interface{}
}

type User struct {
	Cellphone string `json: cellphone`
	Password  string `json: password`
}

func (user *User) CellphoneLogin() interface{} {
	password := []byte(user.Password)
	user.Password = fmt.Sprintf("%x", md5.Sum(password))
	return user
}
