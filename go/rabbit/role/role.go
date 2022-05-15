package role

const (
	PRODUCER = "p"
	CONSUMER = "c"
	ALL      = "a"
)

var Role string

func IsProducer() bool {
	return Role == PRODUCER
}

func IsConsumer() bool {
	return Role == CONSUMER
}
