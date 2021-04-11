package user

import "github.com/gus-messagi/pokedex-api/pokedex-auth-service/pkg/entities"

type Service interface {
	InsertUser(user *entities.User) (*entities.User, error)
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
	return s.repository.CreateUser(user)
}
