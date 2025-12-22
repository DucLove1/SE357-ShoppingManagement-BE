package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OwnerID primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id"`

	Name        string `bson:"name" json:"name"`
	Logo        string `bson:"logo" json:"logo"`
	Cover       string `bson:"cover" json:"cover"`
	Description string `bson:"description" json:"description"`

	PhoneNumber string  `bson:"phone_number" json:"phone_number"`
	Email       string  `bson:"email" json:"email"`
	Address     Address `bson:"address" json:"address"`

	Status             StoreStatus       `bson:"status" json:"status"`
	VerificationStatus StoreVerifyStatus `bson:"verification_status" json:"verification_status"`

	Rating        float64 `bson:"rating" json:"rating"`
	TotalRating   float64 `bson:"total_rating" json:"total_rating"`
	TotalProduct  float64 `bson:"total_product" json:"total_product"`
	TotalOrder    float64 `bson:"total_order" json:"total_order"`
	TotalRevenue  float64 `bson:"total_revenue" json:"total_revenue"`
	FollowerCount int     `bson:"follower_count" json:"follower_count"`

	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
	VerifiedAt time.Time `bson:"verified_at" json:"verified_at"`
}

type StoreStatus string

const (
	StoreActive   StoreStatus = "active"
	StoreInactive StoreStatus = "inactive"
	StoreBanned   StoreStatus = "banned"
	StorePending  StoreStatus = "pending"
)

type StoreVerifyStatus string

const (
	VerifyPending  StoreVerifyStatus = "pending"
	VerifyApproved StoreVerifyStatus = "approved"
	VerifyRejected StoreVerifyStatus = "rejected"
)
