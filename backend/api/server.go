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

	server.router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Next()
	})

	server.router.POST("/user", server.AddUser)
	server.router.POST("/chat", server.CreateChat)
	server.router.GET("/chat/:username", server.GetUserChat)
	server.router.POST("/message", server.CreateMessage)
	server.router.GET("/message", server.GetChatMessages)
	server.router.POST("/qa", server.CreateQA)
	server.router.GET("/qa", server.GetAllQA)
	server.router.DELETE("/qa", server.RemoveQA)
	server.router.PUT("/qa", server.UpdateQA)

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
