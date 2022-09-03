package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rest-api-test/internal/apperrors"
	"rest-api-test/internal/user"
	"rest-api-test/pkg/logging"
)

type storage struct {
	collection *mongo.Collection
	cache      *redis.Client
	logger     *logging.Logger
}

func NewStorage(database *mongo.Database, rdb *redis.Client, collection string, logger *logging.Logger) user.Storage {
	return &storage{
		collection: database.Collection(collection),
		cache:      rdb,
		logger:     logger,
	}
}

func (d *storage) Create(ctx context.Context, user user.User) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to: %v", err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectID to hex: %v", err)
}
func (d *storage) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to convert to objectId: %s", id)
	}

	filterFindOne := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filterFindOne)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return user.User{}, apperrors.ErrNotFound
		}
		return u, fmt.Errorf("failed to find user by id: %s, due to error: %v", id, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user by id: %s, due to error: %v", id, err)
	}

	return u, nil
}
func (d *storage) Update(ctx context.Context, user user.User) error {
	panic("implement me")
}
func (d *storage) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
