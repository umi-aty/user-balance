package repositories

import (
	"log"
	"net/http"
	"userbalance/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user entities.User) entities.User
	IsDuplicateEmail(email string) (db *gorm.DB)
	Login(user entities.User) entities.User
	EmailNotFound(email string, password string) string
}

type userConnection struct {
	db         *gorm.DB
	httpClient *http.Client
}

func NewUserRepo(db *gorm.DB, httpClient *http.Client) UserRepository {
	return &userConnection{
		db:         db,
		httpClient: httpClient,
	}
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("failed to hash password")
	}
	return string(hash)
}

func (connection *userConnection) Register(user entities.User) entities.User {
	user.Password = hashAndSalt([]byte(user.Password))
	connection.db.Save(&user)
	return user
}

func (connection *userConnection) Login(user entities.User) entities.User {
	connection.db.Where("email = ?", user.Email).Take(&user)
	connection.db.Find(&user)
	return user
}

func (connection *userConnection) IsDuplicateEmail(email string) (db *gorm.DB) {
	var user entities.User
	return connection.db.Where("email = ?", email).Take(&user)
}

func (connection *userConnection) EmailNotFound(email string, password string) string {
	var user entities.User
	connection.db.First(&user, "email = ?", email)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "not match"
	} else {
		return "match"
	}
}
