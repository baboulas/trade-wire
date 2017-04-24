package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"database/sql"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Type     string    `json:"type"`
	Password []byte    `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}

// NewUser {u} is an instance of user struct
func NewUser(id uuid.UUID, name, username, t string, password []byte) *User {
	return &User{
		ID:       id,
		Name:     name,
		Username: username,
		Type:     t,
		Password: password,
	}
}

func NewUserLogin(username string, password []byte) *UserLogin {
	return &UserLogin{
		Username: username,
		Password: password,
	}
}

// HashPassword hashes password field from incoming requests
func (u *User) hashPassword() []byte {
	hfp, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hfp
}

func (u *User) Save(db *gorm.DB) {
	p := u.hashPassword()
	db.Table("users").Create(&User{
		u.ID,
		u.Name,
		u.Username,
		u.Type,
		p,
	})
}

func (u *User) Authorize(db *gorm.DB) {
	// u.checkForUser(db)
}

func FetchAllUsers(db *gorm.DB) []User {
	// var results UserResults
	var users []User
	db.Select([]string{"id", "name", "username", "type"}).Find(&users)
	return users
}

// func (ur *UserResults) FetchAll(db *gorm.DB) {
// 	var results []UserResults
// 	db.Table("users").
// }

func (u *User) Update(db *sql.DB, id uuid.UUID, ud *User) {
}

func (u *User) Destroy(db *sql.DB) {
}

func (ul *UserLogin) Auth(db *gorm.DB) map[string]string {
	user, err := ul.checkForUser(db)
	if err != "" {
		return map[string]string{
			"error": err,
		}
	}

	token := ul.generateToken()
	return map[string]string{
		"id":    uuid.UUID.String(user.ID),
		"token": token,
	}

}

func (ul *UserLogin) generateToken() string {
	var mySigningKey = []byte("supersecretkey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func (ul *UserLogin) checkForUser(db *gorm.DB) (*User, string) {
	// var data UserLogin
	var user User
	var err string
	db.Table("users").Where("username = ?", ul.Username).Find(&user)
	fmt.Print(user)
	if user.Name == "" {
		err = "user not found"
	}
	return &user, err
}
