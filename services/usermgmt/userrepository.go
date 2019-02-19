package usermgmt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var uuid = 0

func NewUserFromRequest(req *UserCreateRequest) *UserModel {
	newUser := &UserModel{
		Username: req.Username,
		Password: hashPassword(req.Password),
		ID:       uuid,
	}
	uuid++
	return newUser
}

func hashPassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
