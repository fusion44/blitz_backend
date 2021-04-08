# Raspberry Blitz Backend

This project is in its very early stages of development and may contain bugs, use with caution.

## Basic Setup

### Postgres

A working PostgresSQL instance is required. Update the .env and the config.toml files with the Postgres credentials.

#### Run a Postgres instance with docker
For development purposes a simple docker instance will do:

`docker run --name blitz-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres`

### Execute migrations

Execute these commands from within the root project folder or change the path's accordingly. Only migrate version v4.11.0 was tested thus far, but a newer version should work as well. At some point in the future this step will be redundant.

#### Get migrate executable

In project root:\
`curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz`

#### Migrate up

Make sure you have edited the `.env` file with the postgres instance credential before running this.\
`source .env && ./migrate.linux-amd64 -path db/migrations -database $POSTGRESQL_URL up`

### Run the server

In project root:\ `go run server.go`

Open GraphQL Playground: [http://localhost:8081](http://localhost:8081)

## Develop
### Reset the database

```sh
source ./.env

./migrate.linux-amd64 -path db/migrations -database $POSTGRESQL_URL drop
./migrate.linux-amd64 -path db/migrations -database $POSTGRESQL_URL up
```

### Generate Dataloader code

```sh
cd backend/graph/model
go run github.com/vektah/dataloaden UserLoader string *github.com/raspiblitz-backend/graph/model/model.User
```

### Update generated files on schema change or name change
Run in project root: `go generate ./...`\
Alternatively run this: `go get github.com/99designs/gqlgen/cmd@v0.13.0 && go run github.com/99designs/gqlgen`

## Changelog

See [CHANGELOG](CHANGELOG.md)

## License

[AGPL V3](LICENSE)
