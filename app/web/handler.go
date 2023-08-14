package web

import (
	"continential/app/persist"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	db persist.DBLayer
}

func (h *Handler) GetContents(c *gin.Context) {

	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error for dsn",
		})
		return
	}

	contents, err := h.db.GetAllConnections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, contents)
}

type HandlerInterface interface {
	GetContents(c *gin.Context)
	GetContent(c *gin.Context)
}
