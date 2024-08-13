package player

import (
	"context"
	"fmt"

	"github.com/arq-hex-go/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) Get(id string) (player domain.Player, err error) {

	// Conviertiendo el id a ObjetctID si es necesario
	objetctId, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		return player, err

	}

	err = r.Client.Database("go-l").Collection("players").FindOne(context.Background(), bson.M{"_id": objetctId}).Decode(&player)

	if err != nil {

		if err == mongo.ErrNoDocuments {

			// Manejar el error en caso de que no se encontró el documento
			return player, fmt.Errorf("no se encontró ningún jugador con id %s", id)

		}

	}

	return player, nil

}