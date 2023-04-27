package serve

import (
	"io"
	"net"
	"testing"
)

func BenchmarkNetworks(b *testing.B) {
	unixgramSocketFile, unixgramData := "/tmp/socket.unixgram.bench", make(chan []byte, 1)
	Datagram("unix:"+unixgramSocketFile, unixgramData)
	unixgramSocket, _ := net.Dial("unixgram", unixgramSocketFile)
	defer unixgramSocket.Close()

	udpPort, udpData := ":8000", make(chan []byte, 1)
	Datagram(udpPort, udpData)
	udpSocket, _ := net.Dial("udp", udpPort)
	defer udpSocket.Close()

	tcpPort, tcpData := ":9000", make(chan []byte, 1)
	Stream(tcpPort, tcpData)
	tcpSocket, _ := net.Dial("tcp", tcpPort)
	defer tcpSocket.Close()

	unixSocketFile, unixData := "/tmp/socket.unix.bench", make(chan []byte, 1)
	Stream("unix:"+unixSocketFile, unixData)
	unixSocket, _ := net.Dial("unix", unixSocketFile)
	defer unixSocket.Close()

	bench := func(w io.WriteCloser, data chan []byte, b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if _, err := w.Write([]byte("test data\n")); err != nil {
				b.Fatalf("error writing data: %v", err)
			}
			<-data
		}
	}

	b.ResetTimer()

	b.Run("unixgram", func(b *testing.B) {
		bench(unixgramSocket, unixgramData, b)
	})
	b.Run("unix", func(b *testing.B) {
		bench(unixSocket, unixData, b)
	})
	b.Run("udp", func(b *testing.B) {
		bench(udpSocket, udpData, b)
	})
	b.Run("tcp", func(b *testing.B) {
		bench(tcpSocket, tcpData, b)
	})
}
