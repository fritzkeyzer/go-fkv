build:
	go work sync
	go build ./...

test:
	go work sync
	go build ./...
	go test -v ./tests