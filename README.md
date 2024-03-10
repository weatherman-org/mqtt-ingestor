# Telemetry Server

## Getting started

### Non containerised version

The following dependencies are required to run the non-containerised version of this application: Go, Postgres, and an MQTT broker.

Use Docker to start an MQTT broker locally with: `docker run -d --name emqx -p 18083:18083 -p 1883:1883 emqx:latest`

This will spin up the EMQX container, its dashboard is available on [localhost:18083](localhost:18083) with default credentials as admin, public. Connect to the MQTT broker via `tcp://localhost:1883`.

Use Docker to start and setup Postgres locally with: `make setup_postgres`

This will create a Postgres container. Connect to the database via `postgresql://postgres:password@localhost:5432/telemetry?sslmode=disable`

Start the MQTT subscriber and HTTP server with `go run main.go`.

### Containerised version

The following dependencies are required to run the containerised version of this application: Docker. Use Docker Compose to run the Go application and MQTT broker:

`docker-compose up -d`

## Testing the application

Install CURL (or your preferred HTTP request platform) and/or [mosquitto](https://mosquitto.org). Publish data via the POST endpoint at `http://localhost:8080/publish` like the following to simulate data publishing to MQTT:

`curl -X POST -H "Content-Type: application/json" -d '{"topic": "topic/telemetry"}' http://localhost:8080/publish`

Publishes can be done through the mosquitto client directly:

`mosquitto_pub -h localhost -p 1883 -t topic/telemetry -m "my amazing message"`

## Cleanup

### Non containerised version

- Run the following to remove the EMQX broker: `docker stop emqx && docker rm emqx`
- Run the following to remove the database: `make postgresdown`

### Containerised version

- Run the following to remove the Docker Compose deployments: `docker-compose down`

### Database data removal

- Remove the database data by removing the created `postgres` directory.
