all: build run
build:
	go build -o main ./main.go
run:
	go run ./main.go