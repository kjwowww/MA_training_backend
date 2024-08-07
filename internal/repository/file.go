package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// File represents the file model in the database
type File struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FileName  string             `bson:"file_name"`
	FileSize  int64              `bson:"file_size"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
