package user

import (
	"context"

	"github.com/gus-messagi/pokedex-api/pokedex-auth-service/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {
	user.ID = primitive.NewObjectID()

	password := []byte(user.Password)

	hashPassword, hashErr := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if hashErr != nil {
		return nil, hashErr
	}

	user.Password = string(hashPassword)

	_, err := r.Collection.InsertOne(context.Background(), user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
