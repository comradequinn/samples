package main

import (
	"flag"
	"log"
	"socket/serve"
)

func main() {
	addr := flag.String("address", ":9000", "the address to listen. either an ip host:port or a unix domain socket expressed as unix:path/socket.example")
	datagram := flag.Bool("datagram", false, "whether to expose a datagram or streamed interface")

	flag.Parse()

	data := make(chan []byte)

	serveFunc := serve.Stream

	if *datagram {
		serveFunc = serve.Datagram
	}

	serveFunc(*addr, data)

	for d := range data {
		log.Println(*addr, "received:", string(d))
	}
}
