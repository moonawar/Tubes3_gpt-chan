--name GetAllQA :many
SELECT * FROM QA

--name GetQAById :one
SELECT * FROM QA WHERE qa_id = $1

--name CreateQA :one
INSERT INTO QA (question, answer) VALUES ($1, $2)

--name RemoveQA :one
DELETE FROM QA WHERE qa_id = $1

--name UpdateQA :one
UPDATE QA SET question = $2, answer = $3 WHERE qa_id = $1