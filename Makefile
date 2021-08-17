.PHONY: clean test security build run

APP_NAME = avitochat
BUILD_DIR = build
MIGRATIONS_FOLDER = internal/database/migrations
DATABASE_URL = postgres://postgres:password@localhost:5432/postgres?sslmode=disable

clean:
	rm -rf ./build

linter:
	golangci-lint run

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) cmd/avitochat/main.go

run: linter build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

docker.build:
	docker build -t avitochat .

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.run: docker.network docker.postgres docker.net_http migrate.up

docker.stop:
	docker stop avitochat avitochat-postgres

docker.net_http: docker.build
	docker run --rm -d \
		--name avitochat \
		--network dev-network \
		-p 9000:9000 \
		avitochat

docker.postgres:
	docker run --rm -d \
		--name avitochat-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres:13.3-alpine
