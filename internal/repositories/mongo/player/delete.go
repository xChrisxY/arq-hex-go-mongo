package player

import (
	"context"
	"errors"
	"log/slog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/jairogloz/go-l/pkg/domain"
)

var (
	ErrDeletePlayer = errors.New("error deleting player")
)

func (r Repository) Delete(id string) (err error) {

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		return err

	}

	collection := r.Client.Database("go-l").Collection("players")
	deleteResult, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectId})

	if err != nil {

		slog.Error("deleting player: ", slog.Any("mongodb", err))
		return ErrDeletePlayer

	}

	if deleteResult.DeletedCount == 0 {

		slog.Error("player not found")
		return domain.ErrNotFound

	}

	return nil

}