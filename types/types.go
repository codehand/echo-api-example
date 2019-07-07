package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product ..
type Product struct {
	ID                    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                  string        `json:"name" bson:"name" nosql:"name" validate:"required"`
	ImageClosed           string        `json:"image_closed" bson:"image_closed" validate:"required"`
	ImageOpen             string        `json:"image_open" bson:"image_open" validate:"required"`
	Description           string        `json:"description" bson:"description" validate:"required"`
	Story                 string        `json:"story" bson:"story" validate:"required"`
	SourcingValues        []string      `json:"sourcing_values" bson:"sourcing_values" validate:"required"`
	Ingredients           []string      `json:"ingredients" bson:"ingredients" validate:"required"`
	AllergyInfo           string        `json:"allergy_info" bson:"allergy_info" validate:"required"`
	DietaryCertifications string        `json:"dietary_certifications" bson:"dietary_certifications" validate:"required"`
	ProductID             string        `json:"product_id" bson:"product_id"`
	CreatedAt             time.Time     `json:"-" bson:"created_at"`
	UpdatedAt             time.Time     `json:"-" bson:"updated_at"`
	DeletedAt             *time.Time    `json:"-" bson:"deleted_at"`
}

// ProductUpdate ..
type ProductUpdate struct {
	ID                    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                  string        `json:"name" bson:"name" nosql:"name" validate:"-"`
	ImageClosed           string        `json:"image_closed" bson:"image_closed" nosql:"image_closed" validate:"-"`
	ImageOpen             string        `json:"image_open" bson:"image_open" nosql:"image_open" validate:"-"`
	Description           string        `json:"description" bson:"description" nosql:"description" validate:"-"`
	Story                 string        `json:"story" bson:"story" nosql:"story" validate:"-"`
	SourcingValues        []string      `json:"sourcing_values" bson:"sourcing_values" nosql:"sourcing_values" validate:"-"`
	Ingredients           []string      `json:"ingredients" bson:"ingredients" nosql:"ingredients" validate:"-"`
	AllergyInfo           string        `json:"allergy_info" bson:"allergy_info" nosql:"allergy_info" validate:"-"`
	DietaryCertifications string        `json:"dietary_certifications" bson:"dietary_certifications" nosql:"dietary_certifications" validate:"-"`
	ProductID             string        `json:"product_id" bson:"product_id"`
	CreatedAt             time.Time     `json:"-" bson:"created_at"`
	UpdatedAt             time.Time     `json:"-" bson:"updated_at"`
	DeletedAt             *time.Time    `json:"-" bson:"deleted_at"`
}
