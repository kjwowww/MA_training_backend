package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the user model in the database
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// ResetPasswordRequest represents the request body for resetting a password
type ResetPasswordRequest struct {
	Password string `json:"password"`
}
