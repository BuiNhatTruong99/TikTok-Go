DB_URL=postgresql://root:secret@localhost:5432/tiktok?sslmode=disable

postgres:
	docker run --name postgres16 -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root tiktok

dropdb:
	docker exec -it postgres16 dropdb tiktok

new-migration:
	migrate create -ext sql -dir migrations -seq $(name)

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

mock:
	mockgen --build_flags=--mod=mod -destination internal/comment/mock/pg_mock.go -source internal/comment/interface.go

test:
	go test -v -cover -short ./...

run:
	go run main.go


.PHONY: postgres createdb dropdb migrateup migratedown sqlc mock test run

