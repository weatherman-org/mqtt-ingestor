# MQTT Ingestor

## Getting started

### Non containerised version

The following dependencies are required to run the non-containerised version of this application: Go, an MQTT broker, CURL (or your preferred HTTP request platform) and [mosquitto](https://mosquitto.org). Use Docker to start an MQTT broker locally with:

`docker run -d --name emqx -p 18083:18083 -p 1883:1883 emqx:latest`

This will spin up the EMQX container, its dashboard is available on [localhost:18083](localhost:18083) with default credentials as admin, public. Connect to the MQTT broker via `tcp://localhost:1883`.

Start the MQTT subscriber and HTTP server with `go run main.go`.

### Containerised version

The following dependencies are required to run the containerised version of this application: Docker, CURL (or your preferred HTTP request platform) and [mosquitto](https://mosquitto.org). Use Docker Compose to run the Go application and MQTT broker:

`docker-compose up`

## Testing the application

Publish data via the POST endpoint at `http://localhost:8080/publish` like the following to simulate data publishing to MQTT:

`curl -X POST -H "Content-Type: application/json" -d '{"topic": "topic/telemetry"}' http://localhost:8080/publish`

Publishes can be done through the mosquitto client directly:

`mosquitto_pub -h localhost -p 1883 -t topic/telemetry -m "my amazing message"`
