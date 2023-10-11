postgres13:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13.9 

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres13 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package=mockdb -destination=db/mock/store_mock.go github.com/abhisheksatish1999/simplebank/db/sqlc Store

.PHONY: postgres13 createdb dropdb migrateup migratedown sqlc test server mock