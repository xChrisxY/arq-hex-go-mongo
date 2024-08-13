package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/core"
)

func (h Handler) DeletePlayer(c *gin.Context) {

	playerIdParams := c.Param("id")
	err := h.PlayerService.Delete(playerIdParams)

	if err != nil { core.RespondError(c, err); return }

	c.Status(http.StatusNoContent)

}