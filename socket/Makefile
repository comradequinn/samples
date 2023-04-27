.PHONY: *

PORT=9000
SOCK=/tmp/socket.example

build: 
	@-rm -rf bin 2> /dev/null
	@mkdir bin
	@go build -o bin/socket main.go

stop :
	@-pkill socket

udp : build stop
	@bin/socket --datagram --address ":${PORT}" &
	@echo "enter data:"
	@-nc -u localhost ${PORT}

unixgram : build stop
	@bin/socket --datagram --address "unix:${SOCK}" &
	@echo "enter data:"
	@nc -uU ${SOCK}

tcp : build stop
	@bin/socket --address ":${PORT}" &
	@echo "enter data:"
	@-nc localhost ${PORT}

unix : build stop
	@bin/socket --address "unix:${SOCK}" &
	@echo "enter data:"
	@-nc -U ${SOCK}

bench : stop
	@go test -run=^$$ -bench . -benchmem ./...
