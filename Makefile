test:
	go test -v

generate-cover:
	go test -coverprofile cover.out

open-cover:
	go tool cover -html=cover.out

test-cover: generate-cover open-cover

build:
	go build -o bin/main main.go app.go article.go publisher.go category.go

run:
	./bin/main

all: test build run