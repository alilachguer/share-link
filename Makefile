# .env must be populated with the necessary values
include .env

start-db:
	docker-compose up -d db

stop-db:
	docker-compose stop db

connect-to-db:
	docker exec -it --env-file=.env db bash -c "mysql -u ${DB_USER_NAME} -p${DB_PASSWORD}"

db-create:
	sqlx database create --database-url "sqlite://${DATABASE_URL}"

db-drop:
	sqlx database drop --database-url "sqlite://${DATABASE_URL}"

migration-create:
	sqlx migrate add -r init

migrate-up:
	sqlx migrate run --database-url "sqlite://${DATABASE_URL}"

migrate-down:
	sqlx migrate revert --database-url "sqlite://${DATABASE_URL}"

build:
	if [ -f "${BINARY}" ]; then rm ${BINARY}; fi
	go build -o ${BINARY} cmd/main.go

run: build
	./${BINARY}

tests:
	go test -v ./test
