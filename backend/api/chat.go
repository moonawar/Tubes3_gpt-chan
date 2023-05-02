package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Username string `json:"username" uri:"username" binding:"required,min=1"`
}

func (server *Server) CreateChat(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	chat, err := server.query.CreateChat(c, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, chat)
}

func (server *Server) GetUserChat(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	chat, err := server.query.GetUserChat(c, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, chat)
}
