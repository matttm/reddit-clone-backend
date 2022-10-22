# reddit-clone-backend
![Coverage](https://img.shields.io/badge/Coverage-73.9%25-brightgreen)

Development Status: Completed

## Description

This is the new backend for my reddit-clone-backend written in Go. The original backend can be found in matttm/reddit-clone. The development of that though has ceased, being replaced by this.

## Getting Started

### Running

First you will need to have the docker container running, so in a terminal, run
```
docker compose up
```
Then in another terminal. you can start the server.
```
go run server.go
```
**Note**, both commands need to be run in the root of the project.

### Generating data

Once the docker container is running, start the server once to run the migrations, and run the following command to generate posts and persons to the database.

```
go run scripts/generate.go
```

### Testing

To test, run
```
go test -v ./...
```

## Helpful Commands

### Migrate

```
go run github.com/migrate create -ext sql -dir mysql -seq create_users_table
```
This command creates a migration file for the database.

### Generating GraphQL schema

```
go run github.com/99designs/gqlgen generate
```
This command is used to generate a `resolvers.go` from a `schema.graphqls`

## Author

matttm : Matt Maloney
