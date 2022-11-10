package repository

import (
	"context"
	"errors"
	"github/abraxas-365/Leak/internal/leakerrs"
	"github/abraxas-365/Leak/pkg/user/domain/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*CreateUser*/
func (r *repo) CreateUser(newUser models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return errors.New(leakerrs.InternalError)
	}
	return nil
}

/*UpdateUser*/
func (r *repo) UpdateUser(edit models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	if _, err := collection.UpdateByID(ctx, edit.Id, edit); err != nil {
		return errors.New(leakerrs.InternalError)
	}
	return nil
}

/*Get User using key and a value */
func (r *repo) GetUser(key string, value interface{}) (models.User, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	user := new(models.User)
	if err := collection.FindOne(ctx, bson.M{key: value}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return models.User{}, false, nil
		}
		return models.User{}, false, errors.New(leakerrs.InternalError)
	}

	return *user, true, nil
}
