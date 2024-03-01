package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserInput struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
}

type UpdateUserInput struct {
	ID       string `json:"id" binding:"required"`
	UserName string `json:"user_name"`
	Gender   string `json:"gender"`
	Active   bool   `json:"active"`
}

type Users struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserName  string             `bson:"user_name,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Gender    string             `bson:"gender,omitempty"`
	Active    bool               `bson:"active,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}
