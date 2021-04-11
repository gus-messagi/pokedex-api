package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omiempty"`
	Name     string             `json:"name" binding:"required" bson:"title"`
	Password string             `json:"password" binding:"required" bson:"password"`
	Email    string             `json:"email" binding:"required" bson:"email"`
}
