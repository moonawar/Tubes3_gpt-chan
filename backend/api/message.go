package api

import (
	"errors"
	"fmt"
	algo "gpt-chan/algorithm"
	db "gpt-chan/database/models"
	util "gpt-chan/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	SIMILARITY_TRESHOLD = 0.9
)

type GetChatMessagesRequest struct {
	ChatID int32 `form:"chat_id" binding:"required"`
	Limit  int32 `form:"limit"`
	Page   int32 `form:"page"`
}

func (server *Server) GetChatMessages(c *gin.Context) {
	var req GetChatMessagesRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// set default values
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}

	params := db.GetChatMessagesParams{
		ChatID: req.ChatID,
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	user, err := server.query.GetChatMessages(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

type CreateMessageRequest struct {
	ChatID    int32  `json:"chat_id" binding:"required"`
	Question  string `json:"question" binding:"required"`
	Algorithm string `json:"algorithm" binding:"required,oneof=kmp bm"`
}

func (server *Server) CreateMessage(c *gin.Context) {
	var req CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// this is where the string matching is done
	alg := algo.New()
	l_question := strings.ToLower(req.Question) // lowercase question

	// classify the question
	code := alg.Classify(l_question)

	fmt.Println("New request with code: ", code)

	var answer string

	if code/8 >= 1 {
		// contains math expression
		handleMathMessage(l_question, &answer)
		code -= 8
	}

	if code/4 >= 1 {
		// contains date question
		handleDateMessage(l_question, &answer)
		code -= 4
	}

	// all of the code below is for string matching
	qa_table, err := server.query.GetAllQA(c) // get all QA from database
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var candidates_qa util.PriorityQueue // list of QA id that matches the question or most similar
	queueSize := 5                       // number of candidates to be returned

	for _, qa := range qa_table {
		// string matching
		var match int
		switch req.Algorithm {
		case "kmp":
			match = alg.KMP(l_question, qa.Question)
		case "bm":
			match = alg.BM(l_question, qa.Question)
		}

		var similarity float64
		similarity = 1.0

		if match < 0 {
			similarity = alg.HammingDistance(l_question, qa.Question)
		}

		if candidates_qa.Len() == queueSize {
			if similarity < candidates_qa.PriorityAt(queueSize-1) {
				continue
			}
		}
		candidates_qa.PushVal(qa, similarity)
	}

	if code/2 >= 1 {
		// contains QA add request
		// extract the question and answer
		q := algo.ExtractAddQuestions(l_question)
		a := algo.ExtractAnswers(l_question)

		for i := 0; i < len(q); i++ {
			// check if the question already exists
			var exists bool
			var item_match_id int32
			var item_match_idx int
			for _, candidate := range candidates_qa {
				var match int
				switch req.Algorithm {
				case "kmp":
					match = alg.KMP(q[i], candidate.Value().Question)
				case "bm":
					match = alg.BM(q[i], candidate.Value().Question)
				}

				var similarity float64
				similarity = 1.0

				if match < 0 {
					similarity = alg.HammingDistance(q[i], candidate.Value().Question)
				}

				exists = similarity > SIMILARITY_TRESHOLD
				if exists {
					item_match_id = candidate.Value().QaID
					item_match_idx = candidate.Index()
					break
				}
			}

			if !exists {
				// add the question and answer to the database
				params := db.CreateQAParams{
					Question: q[i],
					Answer:   a[i],
				}

				_, err := server.query.CreateQA(c, params)
				if err != nil {
					c.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				answer += "Pertanyaan: `" + q[i] + "` dengan jawaban `" + a[i] + "` berhasil ditambahkan ke database.\n"
			} else {
				params := db.UpdateQAParams{
					QaID:     item_match_id,
					Question: q[i],
					Answer:   a[i],
				}

				_, err := server.query.UpdateQA(c, params)
				if err != nil {
					c.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				answer += "Pertanyaan: `" + q[i] + "` sudah ada! Jawaban diupdate menjadi `" + a[i] + "`.\n"
				candidates_qa.RemoveAt(item_match_idx)
			}
		}
		code -= 2
	}

	if code/1 >= 1 {
		// contains QA delete request
		// extract the question and answer
		q := algo.ExtractDeleteQuestions(l_question)

		for i := 0; i < len(q); i++ {
			// check if the question already exists
			var exists bool
			var item_match_id int32
			var item_match_idx int
			for _, candidate := range candidates_qa {
				var match int
				switch req.Algorithm {
				case "kmp":
					match = alg.KMP(q[i], candidate.Value().Question)
				case "bm":
					match = alg.BM(q[i], candidate.Value().Question)
				}

				var similarity float64
				similarity = 1.0

				if match < 0 {
					similarity = alg.HammingDistance(q[i], candidate.Value().Question)
				}

				exists = similarity > SIMILARITY_TRESHOLD
				if exists {
					item_match_id = candidate.Value().QaID
					item_match_idx = candidate.Index()
					break
				}
			}

			if !exists {
				answer += "Pertanyaan: `" + q[i] + "` tidak ditemukan di database.\n"
			} else {
				_, err := server.query.RemoveQA(c, item_match_id)
				if err != nil {
					c.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				answer += "Pertanyaan: `" + q[i] + "` berhasil dihapus dari database.\n"
				candidates_qa.RemoveAt(item_match_idx)
			}
		}
		code -= 1
	}

	// code must be 0
	if code != 0 {
		err := errors.New("classifier error")
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} else {
		for candidates_qa.Len() > 0 && candidates_qa.PriorityAt(0) > SIMILARITY_TRESHOLD {
			qa := candidates_qa.PopVal()
			answer += qa.Answer + "\n"
		}
	}

	if answer == "" {
		answer = "Maaf, saya tidak mengerti pertanyaan Anda. Silakan coba lagi."
	}

	msg_params := db.CreateMessageParams{
		ChatID:   req.ChatID,
		Question: req.Question,
		Answer:   answer,
	}

	msg, err := server.query.CreateMessage(c, msg_params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, msg)
}

func handleMathMessage(input string, output *string) {
	// contains math question
	// extract the math expression
	math_exps := algo.ExtractMathExps(input)
	alg := algo.New()
	for _, exp := range math_exps {
		// solve the math expression
		res, err := alg.SolveMath(exp)
		res_str := strconv.FormatFloat(res, 'f', 2, 64)
		if err == nil { // success
			*output += "Hasil dari " + exp + " adalah " + res_str + "\n"
		} else { // error
			*output += "Tidak dapat menyelesaikan " + exp + " karena kesalahan sintaks.\n"
			fmt.Println(err)
		}
	}
}

func handleDateMessage(input string, output *string) {
	// contains date question
	// extract the date expression
	dates := algo.ExtractDates(input)
	fmt.Println(dates)
	for _, date := range dates {
		day := algo.DateToDay(date)
		if day == "" {
			*output += "Tanggal " + date + " tidak valid.\n"
		} else {
			*output += "Tanggal " + date + " adalah hari " + algo.DateToDay(date) + ".\n"
		}
	}
}
