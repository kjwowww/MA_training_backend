package model

import "time"

type File struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	FileName   string    `json:"file_name" bson:"file_name"`
	FileSize   int64     `json:"file_size" bson:"file_size"`
	CreateDate time.Time `json:"create_date" bson:"create_date"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
	CreatedBy  string    `json:"created_by" bson:"created_by"`
}
