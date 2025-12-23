package repo

import (
	"context"
	"time"

	"github.com/DucLove1/SE357-ShoppingManagement-BE/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProvinceRepository interface {
	// Province CRUD
	Create(ctx context.Context, province *model.Province) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*model.Province, error)
	GetAll(ctx context.Context) ([]model.Province, error)
	Update(ctx context.Context, province *model.Province) error
	Delete(ctx context.Context, id primitive.ObjectID) error

	// Ward CRUD
	AddWard(ctx context.Context, provinceID primitive.ObjectID, ward *model.Ward) error
	GetWardByID(ctx context.Context, provinceID primitive.ObjectID, wardID string) (*model.Ward, error)
	GetWardsByProvinceID(ctx context.Context, provinceID primitive.ObjectID) ([]model.Ward, error)
	UpdateWard(ctx context.Context, provinceID primitive.ObjectID, ward *model.Ward) error
	DeleteWard(ctx context.Context, provinceID primitive.ObjectID, wardID string) error
}

type provinceRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewProvinceRepository(db *mongo.Database) ProvinceRepository {
	return &provinceRepository{
		db:         db,
		collection: db.Collection("provinces"),
	}
}

// ==================== Province CRUD ====================

func (r *provinceRepository) Create(ctx context.Context, province *model.Province) error {
	province.ID = primitive.NewObjectID()

	if province.Wards == nil {
		province.Wards = []model.Ward{}
	}

	_, err := r.collection.InsertOne(ctx, province)
	return err
}

func (r *provinceRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Province, error) {
	var province model.Province
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&province)
	if err != nil {
		return nil, err
	}
	return &province, nil
}

func (r *provinceRepository) GetAll(ctx context.Context) ([]model.Province, error) {
	var provinces []model.Province
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r *provinceRepository) Update(ctx context.Context, province *model.Province) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": province.ID},
		bson.M{
			"$set": bson.M{
				"name": province.Name,
			},
		},
	)
	return err
}

func (r *provinceRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// ==================== Ward CRUD ====================

func (r *provinceRepository) AddWard(ctx context.Context, provinceID primitive.ObjectID, ward *model.Ward) error {
	ward.ID = primitive.NewObjectID().Hex()
	ward.ProvinceID = provinceID

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": provinceID},
		bson.M{
			"$push": bson.M{"wards": ward},
			"$set":  bson.M{"updated_at": time.Now()},
		},
	)
	return err
}

func (r *provinceRepository) GetWardByID(ctx context.Context, provinceID primitive.ObjectID, wardID string) (*model.Ward, error) {
	var province model.Province
	err := r.collection.FindOne(
		ctx,
		bson.M{
			"_id":      provinceID,
			"wards.id": wardID,
		},
	).Decode(&province)

	if err != nil {
		return nil, err
	}

	for _, ward := range province.Wards {
		if ward.ID == wardID {
			return &ward, nil
		}
	}

	return nil, mongo.ErrNoDocuments
}

func (r *provinceRepository) GetWardsByProvinceID(ctx context.Context, provinceID primitive.ObjectID) ([]model.Ward, error) {
	var province model.Province
	err := r.collection.FindOne(ctx, bson.M{"_id": provinceID}).Decode(&province)
	if err != nil {
		return nil, err
	}
	return province.Wards, nil
}

func (r *provinceRepository) UpdateWard(ctx context.Context, provinceID primitive.ObjectID, ward *model.Ward) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{
			"_id":      provinceID,
			"wards.id": ward.ID,
		},
		bson.M{
			"$set": bson.M{
				"wards.$.name": ward.Name,
			},
		},
	)
	return err
}

func (r *provinceRepository) DeleteWard(ctx context.Context, provinceID primitive.ObjectID, wardID string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": provinceID},
		bson.M{
			"$pull": bson.M{"wards": bson.M{"id": wardID}},
		},
	)
	return err
}
