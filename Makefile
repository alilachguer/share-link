# .env must be populated with the necessary values
include .env

start-db:
	docker-compose up -d db

stop-db:
	docker-compose stop db

connect-to-db:
	docker exec -it --env-file=.env db bash -c "mysql -u ${DB_USER_NAME} -p${DB_PASSWORD}"

create-migration:
	sqlx migrate add -r init

migrate-up:
	sqlx migrate run --database-url "mysql://${DB_USER_NAME}:${DB_PASSWORD}@localhost:3306/${DB_DATABASE_NAME}"

migrate-down:
	sqlx migrate revert --database-url "mysql://${DB_USER_NAME}:${DB_PASSWORD}@localhost:3306/${DB_DATABASE_NAME}"

build:
	if [ -f "${BINARY}" ]; then rm ${BINARY}; fi
	go build -o ${BINARY} main.go

run: build
	./${BINARY}
