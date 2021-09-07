package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultSchema struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedBy primitive.ObjectID `json:"createdBy" bson:"createdBy,omitempty"`
	UpdatedBy primitive.ObjectID `json:"updatedBy" bson:"updatedBy,omitempty"`
	DeletedBy primitive.ObjectID `json:"deletedBy" bson:"deletedBy,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	DeletedAt time.Time          `json:"deletedAt" bson:"deletedAt"`
}

type Pagination struct {
	Data    []interface{} `json:"data" bson:"data"`
	Page    int64         `json:"page" bson:"page"`
	PerPage int64         `json:"perPage" bson:"perPage"`
	Total   int64         `json:"total" bson:"total"`
}
