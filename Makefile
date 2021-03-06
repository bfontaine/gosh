TARGET = gosh

all: $(TARGET)

deps:
	go get -d -t -v ./...

$(TARGET): deps
	go build .

check: deps
	go test -v ./...
