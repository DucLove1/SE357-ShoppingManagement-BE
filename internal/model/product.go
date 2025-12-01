package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	CategoryID primitive.ObjectID `bson:"category_id" json:"category_id"`
	StoreID    primitive.ObjectID `bson:"store_id" json:"store_id"`

	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Status      ProductStatus `bson:"status" json:"status"`

	Images    []string `bson:"images" json:"images"`
	Videos    []string `bson:"videos" json:"videos"`
	Thumbnail string   `bson:"thumbnail" json:"thumbnail"`

	Price         float64 `bson:"price" json:"price"`
	StockQuantity int     `bson:"stock_quantity" json:"stock_quantity"`

	SoldCount   int     `bson:"sold_count" json:"sold_count"`
	Rating      float64 `bson:"rating" json:"rating"`
	ReviewCount int     `bson:"review_count" json:"review_count"`

	IsHasVariants bool `bson:"is_has_variants" json:"is_has_variants"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type ProductStatus string

const (
	ProductAvailable    ProductStatus = "available"
	ProductUnavailable  ProductStatus = "unavailable"
	ProductOutOfStock   ProductStatus = "out_of_stock"
	ProductPending      ProductStatus = "pending"
	ProductDiscontinued ProductStatus = "discontinued"
)

type ProductVariant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`

	SKU  string `bson:"sku" json:"sku"`
	Name string `bson:"name" json:"name"`

	Attributes []VariantAttribute `bson:"attributes" json:"attributes"`

	Image string `bson:"image,omitempty" json:"image,omitempty"`

	Price         float64 `bson:"price" json:"price"`
	StockQuantity int     `bson:"stock_quantity" json:"stock_quantity"`

	IsActive  bool `bson:"is_active" json:"is_active"`
	IsDefault bool `bson:"is_default" json:"is_default"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type VariantAttribute struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}

type VariantOption struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	VariantID    primitive.ObjectID `bson:"variant_id" json:"variant_id"`
	MedicineName string             `bson:"medicine_name" json:"medicine_name"`
	Dosage       string             `bson:"dosage" json:"dosage"`
	Instructions string             `bson:"instructions" json:"instructions"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
