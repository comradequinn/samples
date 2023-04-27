package serve

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func Datagram(addr string, data chan []byte) {
	var (
		conn net.PacketConn
		err  error
	)

	if strings.HasPrefix(addr, "unix:") {
		socket := strings.TrimPrefix(addr, "unix:")
		_ = os.Remove(socket)
		conn, err = net.ListenPacket("unixgram", socket)
	} else {
		conn, err = net.ListenPacket("udp", addr)
	}

	if err != nil {
		log.Fatalf("error creating datagram listener for [%v]: %v", addr, err)
	}

	go func() {
		defer conn.Close()

		buf := make([]byte, 1024)

		for {
			n, _, err := conn.ReadFrom(buf)

			if err != nil {
				log.Fatalln("error reading data from socket:", err)
			}

			data <- buf[:n]
		}
	}()
}

func Stream(addr string, data chan []byte) {
	var (
		listener net.Listener
		err      error
	)

	if strings.HasPrefix(addr, "unix:") {
		socket := strings.TrimPrefix(addr, "unix:")
		_ = os.Remove(socket)
		listener, err = net.Listen("unix", socket)
	} else {
		listener, err = net.Listen("tcp", addr)
	}

	if err != nil {
		log.Fatalf("error creating stream listener for [%v]: %v", addr, err)
	}

	go func() {
		defer listener.Close()

		for {
			conn, err := listener.Accept()

			if err != nil {
				log.Fatalln("error accepting connection:", err)
			}

			go func(rc io.ReadCloser) {
				defer rc.Close()

				reader := bufio.NewReader(rc)

				for {
					b, err := reader.ReadBytes('\n')

					if err != nil {
						if err == io.EOF {
							break
						}
						log.Fatalln("error reading from connection:", err)
					}

					data <- []byte(b)
				}
			}(conn)
		}
	}()
}
