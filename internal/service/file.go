package service

import (
	"context"
	"ma-backend-training/config"
	"ma-backend-training/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveFileInfo saves the information of the uploaded file
func SaveFileInfo(ctx context.Context, fileName string, fileSize int64) error {
	collection := config.DB.Collection("files")

	file := repository.File{
		ID:        primitive.NewObjectID(),
		FileName:  fileName,
		FileSize:  fileSize,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, file)
	return err
}

// GetFiles retrieves the list of files from the database
func GetFiles(ctx context.Context) ([]repository.File, error) {
	collection := config.DB.Collection("files")

	var files []repository.File
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var file repository.File
		if err := cursor.Decode(&file); err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

// DeleteFile deletes a file from the database
func DeleteFile(ctx context.Context, fileID string) error {
	collection := config.DB.Collection("files")

	objID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
