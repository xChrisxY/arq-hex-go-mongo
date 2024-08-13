package player

import (
	"errors"

	"github.com/arq-hex-go/internal/domain"
)

func (s Service) Get(id string) (player domain.Player, err error) {

	if id == "" {

		return player, errors.New("id is required")

	}	

	player, err = s.Repo.Get(id)

	if err != nil {

		return player, err

	}

	return player, nil

}