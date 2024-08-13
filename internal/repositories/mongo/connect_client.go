package mongo

import (
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(dbURI string) (client *mongo.Client, err error){

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbURI))

	if err != nil {

		return nil, err

	}

	err = client.Ping(ctx, nil)

	if err != nil {

		return nil, err

	}

	return client, nil


}