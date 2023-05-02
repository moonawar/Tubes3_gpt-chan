package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddUserRequest struct {
	Username string `json:"username" binding:"required,min=1"`
}

func (server *Server) AddUser(c *gin.Context) {
	var req AddUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.query.CreateUser(c, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, user)
}
