# RabbitMQ

## docker

`docker run -p 5672:5672 --name rabbit rabbitmq:3`

## Run

Run Producer: `go run main.go -r p simple`

Run Consumer: `go run main.go -r c simple`

Run All: `go run main.go simple`

# Notice

Please keep in mind that this and other tutorials are, well, tutorials. They demonstrate one new concept at a time and may intentionally oversimplify some things and leave out others. For example topics such as connection management, error handling, connection recovery, concurrency and metric collection are largely omitted for the sake of brevity. Such simplified code should not be considered production ready.

## docs

https://www.rabbitmq.com/documentation.html