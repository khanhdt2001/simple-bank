migrateup:
	go get -u -d github.com/golang-migrate/migrate
	migrate -path db/migration -database "postgresql://postgres:v8hlDV0yMAHHlIurYupj@localhost:5434/simplebank?sslmode=disable" -verbose up

migratedown:
	go get -u -d github.com/golang-migrate/migrate
	migrate -path db/migration -database "postgresql://postgres:v8hlDV0yMAHHlIurYupj@localhost:5434/simplebank?sslmode=disable" -verbose down

sqlc-generate:
	sqlc generate

test:
	go test -v -cover ./...

test-coverprofile:
	go test -v ./... -coverprofile=cover.out
	go tool cover -html=cover.out

mock: 
	mockgen -package mockdb  -destination ./db/mock/store.go simple_bank/db/sqlc Store
	
server:
	go run main.go
