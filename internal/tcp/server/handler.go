package tcpserver

import (
	"fmt"
	"io"
	"log/slog"
	"net"
)

type TCPServer struct {
	addr     string
	listener net.Listener
	sig      chan byte

	readBufLimit   uint16
	ClientMessages chan []byte
}

func NewTCPServer(addr string) *TCPServer {
	return &TCPServer{
		addr:           addr,
		readBufLimit:   1024,
		ClientMessages: make(chan []byte),
	}
}

func (s *TCPServer) Start() error {

	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		slog.Error("error", "e", err)
		return err
	}

	s.listener = listener
	defer s.listener.Close()

	go s.accptConnecitonLoop()

	<-s.sig

	return nil
}

// acceoting connections in tcp
func (s *TCPServer) accptConnecitonLoop() {

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			slog.Error("accepting conn err", "e", err)
			continue
		}

		go s.reader(conn)

		fmt.Println("a user connected ", conn.RemoteAddr())
	}
}

func (s *TCPServer) reader(connState net.Conn) {
	defer connState.Close()
	buf := make([]byte, s.readBufLimit)

	//read loop,
	for {
		size, err := connState.Read(buf)

		if err != nil {

			if err == io.EOF {
				fmt.Println("Client will be disconnected")
				break
			}

			slog.Error("unable to read the buffer", "e", err)
			continue
		}

		s.ClientMessages <- buf[:size]

	}
}
