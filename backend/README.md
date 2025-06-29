# Book The Flight

## Prerequisite

- Golang 1.24.x or newer
- PostgreSQL 15
- Docker
- Docker compose

## Quick start

**Run the server using docker-compose**
```
docker-compose down -v && docker-compose up
```

**For local development**

To install all golang dependencies in `go.mod` file
```
go mod download
```
To migrate data
```
go run main.go migration
```
To start the server
```
go run main.go
```
To check the server
```
curl --location 'http://localhost:8080/v1/seat?passenger_id=3b1ea360-3f82-4f59-918e-b7280d64eb76&flight_id=04104ded-8380-4d88-9798-0f28e32a616b'
```

## Project folder structure
- commands: store commands of the app, ex: run rest server, run migration
- controllers: define routes and handlers
- models: define DB models
- repository: define and implement query functions to data sources
- responses: define response structs
- services: define and implement usecase functions