# RabbitMQ

## docker

`docker run -p 5672:5672 --name rabbit rabbitmq:3`

## Run

Run Producer: `go run main.go -r p simple`

Run Consumer: `go run main.go -r c simple`

Run All: `go run main.go simple`
