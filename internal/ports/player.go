package ports

import "github.com/arq-hex-go/internal/domain"

// expone los servicios/funciones de nuestra applicación
type PlayerService interface {

	Create(player domain.Player) (id interface{}, err error)
	Get(id string) (player domain.Player, err error)
	Delete(id string) (err error)

}

type PlayerRepository interface {

	Insert(player domain.Player) (id interface{}, err error)
	Get(id string) (player domain.Player, err error)
	Delete(id string) (err error)

}	