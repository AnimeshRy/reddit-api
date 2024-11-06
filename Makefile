build:
	go build -o bin/reddit-api

run: build
	./bin/reddit-api
