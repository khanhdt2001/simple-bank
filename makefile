migrateup:
	go get -u -d github.com/golang-migrate/migrate
	migrate -path db/migrate -database "postgresql://postgres:v8hlDV0yMAHHlIurYupj@localhost:5434/simplebank?sslmode=disable" -verbose up

migratedown:
	go get -u -d github.com/golang-migrate/migrate
	migrate -path db/migrate -database "postgresql://postgres:v8hlDV0yMAHHlIurYupj@localhost:5434/simplebank?sslmode=disable" -verbose down

sqlc-generate:
	sqlc generate

test:
	go test -v ./... -coverprofile=cover.out
	go tool cover -html=cover.out
