build: 
	@go build -o ./bin/allthinggo ./cmd

run: build
	@./bin/allthinggo

build-producer:
	@go build -o ./bin/producer ./cmd/producer

run-producer: build-producer
	@./bin/producer

build-consumer:
	@go build -o ./bin/consumer ./cmd/consumer

run-consumer: build-consumer
	@./bin/consumer
