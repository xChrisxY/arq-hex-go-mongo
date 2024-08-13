package player

import (
	"fmt"
	"log"
	"time"
	"github.com/arq-hex-go/internal/domain"
)

func (s Service) Create(player domain.Player) (id interface{}, err error){

	// set creation time
	// Save to repo
	// Responder con el id del recurso creado

	player.CreationTime = time.Now().UTC()

	// ======= repo	

	insertedId, err := s.Repo.Insert(player)

	if  err != nil {

		log.Println(err.Error())
		return nil, fmt.Errorf("error creating player: %w", err)

	}

	// ======= repo

	return insertedId, nil 

}