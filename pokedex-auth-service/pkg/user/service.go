package user

import (
	"github.com/gus-messagi/pokedex-api/pokedex-auth-service/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	InsertUser(user *entities.User) (*entities.User, error)
	SignIn(email string, password string) (*entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertUser(user *entities.User) (*entities.User, error) {

	password := []byte(user.Password)

	hashPassword, hashErr := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if hashErr != nil {
		return nil, hashErr
	}

	user.Password = string(hashPassword)

	return s.repository.CreateUser(user)
}

func (s *service) SignIn(email string, password string) (*entities.User, error) {
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if hashErr != nil {
		return nil, err
	}

	return user, err
}
