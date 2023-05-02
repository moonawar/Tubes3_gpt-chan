package api

import (
	db "gpt-chan/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateQARequest struct {
	Question string `json:"question" binding:"required,min=1"`
	Answer   string `json:"answer" binding:"required,min=1"`
}

func (server *Server) CreateQA(c *gin.Context) {
	var req CreateQARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := db.CreateQAParams{
		Question: req.Question,
		Answer:   req.Answer,
	}

	qa, err := server.query.CreateQA(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, qa)
}

func (server *Server) GetAllQA(c *gin.Context) {
	qa, err := server.query.GetAllQA(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, qa)
}

type RemoveQARequest struct {
	QaID int32 `json:"qa_id" binding:"required"`
}

func (server *Server) RemoveQA(c *gin.Context) {
	var req RemoveQARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	qa, err := server.query.RemoveQA(c, req.QaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, qa)
}

type UpdateQARequest struct {
	QaID     int32  `json:"qa_id" binding:"required"`
	Question string `json:"question" binding:"required,min=1"`
	Answer   string `json:"answer" binding:"required,min=1"`
}

func (server *Server) UpdateQA(c *gin.Context) {
	var req UpdateQARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := db.UpdateQAParams{
		QaID:     req.QaID,
		Question: req.Question,
		Answer:   req.Answer,
	}

	qa, err := server.query.UpdateQA(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, qa)
}
