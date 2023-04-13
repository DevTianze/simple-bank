postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qwer1234 -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrate -database "postgres://root:qwer1234@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrate -database "postgres://root:qwer1234@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY:postgres createdb dropdb migrateup migratedown