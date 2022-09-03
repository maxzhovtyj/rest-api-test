package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rest-api-test/internal/config"
	"time"
)

func NewClient(ctx context.Context, cfg *config.Config) (db *mongo.Database, err error) {
	dsn := fmt.Sprintf(
		"mongodb+srv://%s:%s@golangtestcluster.zv1vbyj.mongodb.net/?retryWrites=true&w=majority",
		cfg.Username,
		cfg.Password,
	)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(dsn).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error occurred while connecting to mongodb: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error occurred while ping mongodb: %v", err)
	}

	return client.Database(cfg.Database), nil
}
