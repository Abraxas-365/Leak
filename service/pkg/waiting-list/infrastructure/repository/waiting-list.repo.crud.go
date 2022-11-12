package repository

import (
	"context"
	"errors"
	"github/abraxas-365/Leak/internal/leakerrs"
	"github/abraxas-365/Leak/pkg/waiting-list/domain/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*CreateFollow*/
func (r *repo) CreateFollow(newUser models.Follow) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return errors.New(leakerrs.InternalError)
	}
	return nil
}

/*Get waitinglist using userId  */
func (r *repo) GetWaitingList(userId uuid.UUID) (models.WaitingList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	waitingList := new(models.WaitingList)
	cur, err := collection.Find(ctx, bson.M{"crush": userId})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return models.WaitingList{}, nil
		}
		return models.WaitingList{}, errors.New(leakerrs.InternalError)
	}

	if err = cur.All(ctx, &waitingList); err != nil {
		return models.WaitingList{}, err
	}

	return *waitingList, nil
}

/*Get follow*/
func (r *repo) GetFollow(crushId uuid.UUID, followerId uuid.UUID) (models.Follow, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	follow := models.Follow{}
	if err := collection.FindOne(ctx, bson.M{"crush": crushId, "follower": followerId}).Decode(&follow); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return models.Follow{}, false, nil
		}
		return models.Follow{}, false, errors.New(leakerrs.InternalError)
	}

	return follow, true, nil
}

/*Delete follow*/
func (r *repo) DeleteFollow(crushId uuid.UUID, followerId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	if _, err := collection.DeleteOne(ctx, bson.M{"crush": crushId, "follower": followerId}); err != nil {
		return err
	}
	return nil
}

/*Update Follow*/
func (r *repo) UpdateFollow(follow models.Follow) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	if _, err := collection.UpdateByID(ctx, follow.Id, follow); err != nil {
		return err
	}

	return nil
}
