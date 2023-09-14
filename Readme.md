# Go Challenge - User API

This README provides essential information for setting up and running the "Go Challenge - User API" project, which includes initializing a Go module, creating a PostgreSQL database with Docker, creating a user API and running tests.

## Table of Contents

- [Getting Started](#getting-started)
  - [Initializing a Go Module](#initializing-a-go-module)
  - [Creating a PostgreSQL Database with Docker](#creating-a-postgresql-database-with-docker)
- [Testing](#testing)
- [Enviroment Variables](#environment-variables)

## Getting Started

Follow the steps below to set up and run the "Go Challenge - User API" project.

### Initializing a Go Module

To start a new Go module for your project, use the following commands:

```shell
go mod init example.com/helloworld
go mod tidy
```

This process ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module's packages and dependencies and removes requirements on modules that don't provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

Once you've initialized the Go module, you can build and run your Go application:

```shell
go build
go run main.go
```

---

### Creating a PostgreSQL Database with Docker:

To set up a PostgreSQL database for your project using Docker, follow these steps:

1 - Pull the PostgreSQL Docker image:

```shell
docker pull postgres
```

2 - Create a Docker container for PostgreSQL:

```
docker run -d --name postgresCont -p 5432:5432 -e POSTGRES_PASSWORD=pass123 postgres
```

- “-d” flag specifies that the container should execute in the background.
- “--name” option assigns the container’s name, i.e., “postgresCont”.
- “-p” assigns the port for the container i.e. “5432:5432”.
- “-e POSTGRES_PASSWORD” configures the password to be “pass123”.
- “postgres” is the official Docker image.

```shell
docker exec -it postgresCont bash

psql -h localhost -U postgres

CREATE DATABASE db_name;
```

---

## Testing

Run the next code.

```shell
go test ./pkg/services
```

## Environment Variables

This project relies on certain environment variables for configuration. To set up the necessary environment variables, follow these steps:

1. Create a `.env` file in the root directory of your project.

2. Type the values in the `.env` file with your actual database credentials and other configuration details.

   ```env
   # .env

   DB_HOST=your_database_host
   DB_PORT=your_database_port
   DB_USER=your_database_user
   DB_PASSWORD=your_database_password
   DB_NAME=your_database_name
   SSL_MODE=disable
   ```
