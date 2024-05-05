# Golang ToDo List Project

## It's my first project on Golang

## Project's stack | new skills

* Gin
* Postgres
* Docker
* JWT
* Hashing password
* Swagger
* Docker and Docker Compose

### How to start the project:

1. Make the file .env in root of project:

```
PORT=8000
DB_PORT=5432
DB_HOST=db
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSLMode=disable
```

2. Start the container postgres

```
docker compose up --build -d
```

3. Let's check a browser on URL:
```
http://localhost:8000/swagger/index.html
```