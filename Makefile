postgresup:
	docker run --name telemetry -p 5432:5432 -e PGUSER=postgres -e POSTGRES_PASSWORD=password -d postgres
	@echo "Waiting for PostgreSQL to be ready..."
	@docker run --rm \
		--link telemetry:postgres \
		--entrypoint sh \
		postgres \
		-c 'while ! pg_isready -h "$$POSTGRES_PORT_5432_TCP_ADDR" -p "$$POSTGRES_PORT_5432_TCP_PORT" > /dev/null; do echo "Waiting for PostgreSQL to be ready..."; sleep 1; done'

postgresdown:
	docker stop telemetry
	docker rm telemetry

createdb:
	docker exec -it telemetry createdb --username=postgres --owner=postgres telemetry

dropdb:
	docker exec -it telemetry dropdb telemetry

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/telemetry?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/telemetry?sslmode=disable" -verbose down

build:
	@echo "Deleting the old mqtt-ingestor image..."
	docker image rm papaya147/weatherman-mqtt-ingestor || true
	@echo "Building the backend Docker image..."
	env GOOS=linux CGO_ENABLED=0 go build -o server-app main.go
	docker build -t papaya147/weatherman-mqtt-ingestor:latest .
	@echo "Cleaning up..."

push: build
	@echo "Pushing image to Docker Hub..."
	docker push papaya147/weatherman-mqtt-ingestor:latest