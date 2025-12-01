package model

type Address struct {
	Province string `bson:"province" json:"province"`
	Ward     string `bson:"ward" json:"ward"`
	Detail   string `bson:"detail" json:"detail"`
}
