// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import ()

type Chat struct {
	ChatID   int32  `json:"chat_id"`
	Username string `json:"username"`
}

type Message struct {
	ChatID   int32  `json:"chat_id"`
	No       int32  `json:"no"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type QA struct {
	QaID     int32  `json:"qa_id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type User struct {
	Username string `json:"username"`
}
