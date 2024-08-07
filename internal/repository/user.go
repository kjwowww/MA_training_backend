package repository

import (
	"context"
	"ma-backend-training/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// GetUserByUsername retrieves a user by username
func GetUserByUsername(ctx context.Context, username string) (User, error) {
	var user User
	collection := config.DB.Collection("users")
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}

// CreateUser creates a new user in the database
func CreateUser(ctx context.Context, user User) error {
	collection := config.DB.Collection("users")
	_, err := collection.InsertOne(ctx, user)
	return err
}

// UpdateUser updates a user in the database
func UpdateUser(ctx context.Context, userID primitive.ObjectID, updateData bson.M) error {
	collection := config.DB.Collection("users")
	_, err := collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": updateData})
	return err
}

// DeleteUser deletes a user from the database
func DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	collection := config.DB.Collection("users")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": userID})
	return err
}
