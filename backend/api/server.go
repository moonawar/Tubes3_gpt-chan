package api

import (
	db "gpt-chan/database/models"

	"github.com/gin-gonic/gin"
)

type Server struct {
	query  db.Queries
	router *gin.Engine
}

func NewServer(query *db.Queries) *Server {
	server := &Server{
		query:  *query,
		router: gin.Default(),
	}

	// router.POST("/", server.CreateUser)

	return server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
