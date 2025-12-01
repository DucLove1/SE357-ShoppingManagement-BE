package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"-"` // Ẩn password
	Role        Role               `bson:"role" json:"role"`
	RoleContent RoleContent        `bson:"role_content,omitempty" json:"role_content,omitempty"`
	UserStatus  UserStatus         `bson:"user_status,omitempty" json:"user_status,omitempty"`
	IsVerified  bool               `bson:"is_verified" json:"is_verified"`
	LastLoginAt *time.Time         `bson:"last_login_at,omitempty" json:"last_login_at,omitempty"`
	CreateAt    time.Time          `bson:"create_at" json:"create_at"`
	UpdateAt    time.Time          `bson:"update_at" json:"update_at"`
}

type UserStatus string

const (
	Banned  UserStatus = "banned"
	Active  UserStatus = "active"
	Deleted UserStatus = "deleted"
)

type Role string

const (
	AdminRole  Role = "admin"
	BuyerRole  Role = "buyer"
	SellerRole Role = "seller"
)

type RoleContent struct {
	Buyer  *BuyerRoleContent  `bson:"buyer,omitempty" json:"buyer,omitempty"`
	Seller *SellerRoleContent `bson:"seller,omitempty" json:"seller,omitempty"`
	Admin  *AdminRoleContent  `bson:"admin,omitempty" json:"admin,omitempty"`
}

// Buyer
type BuyerRoleContent struct {
	Avatar      string    `bson:"avatar,omitempty" json:"avatar,omitempty"`
	FullName    string    `bson:"full_name,omitempty" json:"full_name,omitempty"`
	PhoneNumber string    `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Gender      Gender    `bson:"gender,omitempty" json:"gender,omitempty"`
	DateOfBirth time.Time `bson:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`
	Address     *Address  `bson:"address,omitempty" json:"address,omitempty"`
}

// Seller
type SellerRoleContent struct {
	Avatar      string    `bson:"avatar,omitempty" json:"avatar,omitempty"`
	FullName    string    `bson:"full_name,omitempty" json:"full_name,omitempty"`
	PhoneNumber string    `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Gender      Gender    `bson:"gender,omitempty" json:"gender,omitempty"`
	DateOfBirth time.Time `bson:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`
	Address     *Address  `bson:"address,omitempty" json:"address,omitempty"`

	ShopIDs []primitive.ObjectID `bson:"shop_ids,omitempty" json:"shop_ids,omitempty"`

	// CMND/CCCD để xác minh
	IdentityCard       string     `bson:"identity_card,omitempty" json:"identity_card,omitempty"`
	IdentityCardImages []string   `bson:"identity_card_images,omitempty" json:"identity_card_images,omitempty"` // Ảnh CMND
	IsVerified         bool       `bson:"is_verified" json:"is_verified"`
	VerifiedAt         *time.Time `bson:"verified_at,omitempty" json:"verified_at,omitempty"`
}

// Admin
type AdminRoleContent struct {
	FullName    string             `bson:"full_name" json:"full_name"`
	Permissions []Permission       `bson:"permissions" json:"permissions"`
	CreateAt    time.Time          `bson:"create_at" json:"create_at"`
	CreateBy    primitive.ObjectID `bson:"create_by,omitempty" json:"create_by,omitempty"`
}

type Permission string

const (
	PermissionUserManage    Permission = "user:manage"
	PermissionProductManage Permission = "product:manage"
	PermissionOrderManage   Permission = "order:manage"
	PermissionShopManage    Permission = "shop:manage"
	PermissionReportView    Permission = "report:view"
	PermissionSettingManage Permission = "setting:manage"
)

// Enums
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)
