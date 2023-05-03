include app.env
export

postgres:
	docker run --name gpt-chan -p 5432:5432 \
	-e POSTGRES_PASSWORD=${DB_PASS} \
	-e POSTGRES_USER=${DB_USER} \
	-e POSTGRES_DB=${DB_NAME} \
	-d postgres

createdb:
	docker exec -it gpt-chan createdb -U ${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it gpt-chan dropdb -U ${DB_USER} ${DB_NAME}

migrateup:
	migrate -path database/migrations -database "postgresql://${DB_USER}:${DB_PASS}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://${DB_USER}:${DB_PASS}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down

migrateversion:
	migrate -path database/migrations -database "postgresql://${DB_USER}:${DB_PASS}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose version

sqlc:
	cd database && docker run --rm -v "$(WORKING_DIR):/src" -w /src kjconroy/sqlc generate

server:
	go run main.go

image:
	docker build -t gpt-chan:latest .

run:
	docker run --name gpt-chan --network app-network -p 8080:8080 \
	-e GIN_MODE=release \
	-e DB_SOURCE="postgresql://${DB_USER}:${DB_PASS}@postgres:5432/${DB_NAME}?sslmode=disable" \
	gpt-chan:latest

.PHONY: postgres createdb dropdb migrateup migratedown sqlc