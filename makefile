.PHONY: postgres createdb dropdb migrateup migratedown test run build envup envdown sleep redisup

postgres:
	docker run --name lms-postgres --rm -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -d postgres

redisup:
	docker run --name lms-redis --rm -p 6379:6379 -d redis

redisdown:
	docker stop lms-redis

createdb:
	docker exec -it lms-postgres createdb --username=postgres --owner=postgres lms

dropdb:
	docker exec -it lms-postgres dropdb --username=postgres lms
	docker stop lms-postgres

migrateup:
	migrate -path database/migrations -database "postgresql://postgres:123456@localhost/lms?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://postgres:123456@localhost/lms?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

run:
	go run main.go

build:
	go build -o /bin/main main.go

envup: postgres sleep-5 createdb sleep-5 migrateup sleep-5 redisup

envdown: migratedown sleep-5 dropdb sleep-5 redisdown

envupw: postgres timeout-5 createdb timeout-5 migrateup timeout-5 redisup

envdownw: migratedown timeout-5 dropdb timeout-5 redisdown

sleep-%:
	sleep $(@:sleep-%=%)

timeout-%:
	timeout $(@:timeout-%=%)