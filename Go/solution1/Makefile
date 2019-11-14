PROJECTNAME=bookstore
GOFILES=$(wildcard *.go)

build:
	go build -o ./$(PROJECTNAME) $(GOFILES)

test:
	go test ./
	go test -bench=. -benchmem