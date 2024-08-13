package player

import (
	"context"
	"fmt"
	"log"
	"github.com/arq-hex-go/internal/domain"
)

func (r Repository) Insert(player domain.Player) (id interface{}, err error){

	collection := r.Client.Database("go-l").Collection("players")

	insertResult, err := collection.InsertOne(context.Background(), player)

	if err != nil { log.Println(err); return nil, fmt.Errorf("error inserting player: %w", err) }

	return insertResult.InsertedID, nil

} 