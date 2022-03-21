PSQL_DNS := "user=postgresu password=postgrespassword dbname=postgresdb sslmode=disable"

up:
	docker-compose up -d

down:
	docker-compose down

generate:
	go generate ./...

migrate:
	docker-compose exec -T app go run db/migration/migration.go

initdata:
	docker-compose exec -T db psql ${PSQL_DNS} -f tmp/sample/init.sql