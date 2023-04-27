postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qwer1234 -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrate -database "postgres://root:qwer1234@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateupone:
	migrate -path db/migrate -database "postgres://root:qwer1234@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrate -database "postgres://root:qwer1234@localhost:5432/simple_bank?sslmode=disable" -verbose down -all

migratedownone:
	migrate -path db/migrate -database "postgres://root:qwer1234@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

dockersqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/DevTianze/simple-bank/db/sqlc Store

.PHONY:postgres createdb dropdb migrateup migratedown sqlc test server mock migrateupone migrateupdownone