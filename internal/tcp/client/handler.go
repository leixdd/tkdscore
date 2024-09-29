package tcpclient

import (
	"fmt"
	"io"
	"log/slog"
	"net"
)

type TCPClient struct {
	serverAddr     string
	ConnState      net.Conn
	readBufLimit   uint16
	ServerMessages chan []byte
}

func NewTCPClient(serverAddr string) *TCPClient {
	return &TCPClient{
		serverAddr:     serverAddr,
		readBufLimit:   1024,
		ServerMessages: make(chan []byte, 1024),
	}
}

func (state *TCPClient) Connect() {

	conn, err := net.Dial("tcp", state.serverAddr)

	if err != nil {
		slog.Error("unable to connect to server ", "err", err)
		panic(err)
	}

	state.ConnState = conn

	fmt.Println("connected to server: ", state.serverAddr)

}

func (state *TCPClient) Writer(message []byte) error {
	_, err := state.ConnState.Write(message)

	return err
}

func (state *TCPClient) Reader() {
	defer state.ConnState.Close()
	buf := make([]byte, state.readBufLimit)

	//read loop,
	for {
		size, err := state.ConnState.Read(buf)

		if err != nil {

			if err == io.EOF {
				fmt.Println("connected to server")
				break
			}

			slog.Error("unable to read the buffer", "e", err)
			continue
		}

		state.ServerMessages <- buf[:size]

	}
}
