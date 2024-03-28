# Golang ToDo List Project

## It's my first project on Golang

## Project's stack | new skills

* Gin
* Postgres
* Docker
* JWT
* Hashing password
* Swagger

### How to start the project:

1. Make the file .env in root of project:

```
PORT=8000
DB_PORT=5436
DB_HOST=localhost
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSLMode=disable
```

2. Start the container postgres

```
docker run --name=todo-db -e POSTGRES_PASSWORD=postgres -p 5436:5432 --rm -d postgres
```

3. Do migrate

```
migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up
```

4. Finally, you can start the project

```
go run core/main.go
```