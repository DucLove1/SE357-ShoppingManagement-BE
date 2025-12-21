package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"userId" json:"userId"`

	Items []CartItem `bson:"items" json:"items"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type CartItem struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ProductID        string             `bson:"productId" json:"productId"`
	ProductVariantID string             `bson:"productVariantID" json:"productVariantID"`
	StoreID          string             `bson:"storeId" json:"storeId"`

	StoreName string `bson:"storeName" json:"storeName"`
	Name      string `bson:"name" json:"name"`
	SKU       string `bson:"sku" json:"sku"`

	AttributeGroup AttributeGroup `bson:"attributeGroup" json:"attributeGroup"`
	Quantity       int            `bson:"quantity" json:"quantity"`
	Image          string         `bson:"image" json:"image"`
}
