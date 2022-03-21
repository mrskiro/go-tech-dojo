up:
	docker-compose up -d

down:
	docker-compose down

generate:
	go generate ./...

migrate:
	docker-compose exec -T app go run db/migration/migration.go
