package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/core"
)

func (h Handler) GetPlayer(c *gin.Context) {

	// traducir el request
	// int: validation
	// consumir un servicio
	// traducir el response

	playerIdParams := c.Param("id")
	
	player, err := h.PlayerService.Get(playerIdParams)

	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, player)	

}