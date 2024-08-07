package service

import (
	"context"
	"ma-backend-training/config"
	"ma-backend-training/internal/repository"
	"ma-backend-training/internal/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateUser updates a user in the database
func UpdateUser(ctx context.Context, userID string, firstName, lastName, username string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{
		"first_name": firstName,
		"last_name":  lastName,
		"username":   username,
		"updated_at": time.Now(),
	}

	return repository.UpdateUser(ctx, objID, update)
}

// DeleteUser deletes a user from the database
func DeleteUser(ctx context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	return repository.DeleteUser(ctx, objID)
}

// ResetUserPassword resets a user's password
func ResetUserPassword(ctx context.Context, userID string, newPassword string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	update := bson.M{
		"password":   hashedPassword,
		"updated_at": time.Now(),
	}

	return repository.UpdateUser(ctx, objID, update)
}

// GetUsers retrieves a list of users from the database
func GetUsers(ctx context.Context, page, rows int, keyword string) ([]repository.User, error) {
	collection := config.DB.Collection("users")

	filter := bson.M{}
	if keyword != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"first_name": bson.M{"$regex": keyword, "$options": "i"}},
				{"last_name": bson.M{"$regex": keyword, "$options": "i"}},
				{"username": bson.M{"$regex": keyword, "$options": "i"}},
			},
		}
	}

	skip := int64((page - 1) * rows)
	limit := int64(rows)

	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	cursor, err := collection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []repository.User
	for cursor.Next(ctx) {
		var user repository.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
