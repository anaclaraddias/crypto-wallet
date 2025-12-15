run:
	go run ./src

test:
	go test -v -race ./... -covermode=atomic