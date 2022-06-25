# reddit-clone-backend

## Description

This is the new backend for my reddit-clone-backend written in Go. The original backend can be found in matttm/reddit-clone. The development of that though has ceased, being replaced by this.

## Getting Started

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
