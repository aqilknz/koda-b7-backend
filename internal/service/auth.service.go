package service

import (
	"regexp"

	"github.com/aqilknz/koda-b7-backend/internal/dto"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService struct {
	users     []User
	idCounter int
}

func NewUserService() *UserService {
	return &UserService{
		users:     []User{},
		idCounter: 1,
	}
}

func (s *UserService) ValidateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

func (s *UserService) RegisterUser(input dto.UserRequest) (User, error) {
	newUser := User{
		Id:       s.idCounter,
		Email:    input.Email,
		Password: input.Password,
	}
	s.users = append(s.users, newUser)
	s.idCounter++
	return newUser, nil
}

func (s *UserService) LoginUser(input dto.UserRequest) (User, bool) {
	for _, usr := range s.users {
		if usr.Email == input.Email && usr.Password == input.Password {
			return usr, true
		}
	}
	return User{}, false
}
