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