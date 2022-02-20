package services

import (
	"log"
	"userbalance/entities"
	"userbalance/repositories"
	"userbalance/services/request"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	Register(user request.RegisterRequest) entities.User
	IsDuplicateEmail(email string) bool
	Login(user request.LoginRequest) entities.User
	EmailNotFound(email string, password string) string
}

type authService struct {
	user repositories.UserRepository
}

func NewAuthService(user repositories.UserRepository) AuthService {
	return &authService{
		user: user,
	}
}

func (service *authService) Register(user request.RegisterRequest) entities.User {
	userCreate := entities.User{}

	err := smapping.FillStruct(&userCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.user.Register(userCreate)
	return res
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.user.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *authService) Login(user request.LoginRequest) entities.User {
	userLogin := entities.User{}

	err := smapping.FillStruct(&userLogin, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.user.Login(userLogin)
	return res
}

func (service *authService) EmailNotFound(email string, password string) string {
	res := service.user.EmailNotFound(email, password)
	return res
}
