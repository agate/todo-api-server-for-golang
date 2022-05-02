# TODO API Server For Golang

## Dependencies

* Docker
* Docker Compose

## How to start

`docker-compose up`

## How to use

* List all the todos
  `GET localhost:8080/api/todos`
* Add a new todo
  `POST localhost:8080/api/todos`
  `title=xxxx`
* Mark a todo as completed
  `POST localhost:8080/api/todos/:id/completed`
* Mark a todo as incompleted
  `POST localhost:8080/api/todos/:id/incompleted`
