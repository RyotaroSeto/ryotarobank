
DB_URL=postgresql://postgres:postgres@localhost:5432/ryotaro_bank?sslmode=disable

initpostgres:
	docker run --name postgres --network bank-ryotaro_bank -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14.5-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres ryotaro_bank

postgres:
	docker exec -it postgres /bin/bash

su:
	su postgres

psql:
	psql ryotaro_bank

dropdb:
	docker exec -it postgres dropdb ryotaro_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

#たびたびquerier.goに追加されたものを追加する
mock:
	mockgen -package mockdb -source=db/sqlc/store.go -destination db/mock/store.go Store

.PHONY: postgres createdb dropdb  migrateup migratedown migrateup1 migratedown1 sqlc
