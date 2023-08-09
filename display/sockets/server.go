package sockets

import (
	"errors"
	"log"
	"net"
)

type Message struct {
	From    string
	Payload []byte
}

type Server struct {
	ListenAddr     string
	Listener       net.Listener
	Lock           chan struct{}
	MessageChannel chan Message
	Peer           net.Conn
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr:     listenAddr,
		Lock:           make(chan struct{}),
		MessageChannel: make(chan Message, 10),
		Peer:           nil,
	}
}

func (s *Server) Start() error {
	log.Println("Starting server on: ", s.ListenAddr)
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.Listener = ln

	go s.AcceptLoop()

	<-s.Lock
	close(s.MessageChannel)

	return nil
}

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		log.Println("new connection:", conn.RemoteAddr())

		s.Peer = conn
		// go s.ReadLoop(conn)
	}
}

// func (s *Server) ReadLoop(conn net.Conn) {
// 	defer conn.Close()
// 	buffer := make([]byte, 1024)

// 	for {
// 		n, err := conn.Read(buffer)
// 		if err != nil {
// 			log.Println("read error:", err)
// 			return
// 		}

// 		s.MessageChannel <- Message{
// 			From:    s.Peer.RemoteAddr().String(),
// 			Payload: buffer[:n],
// 		}

// 		conn.Write([]byte("Ok"))
// 	}
// }

func (s *Server) Broadcast(message []byte) ([]byte, error) {
	if s.Peer == nil {
		return nil, errors.New("no peer connected")
	}
	if _, err := s.Peer.Write(message); err != nil {
		s.Peer.Close()
		return nil, errors.New("error writing to connection, close at " + s.Peer.RemoteAddr().String() + ", error: " + err.Error())
	}

	buffer := make([]byte, 1024)
	n, err := s.Peer.Read(buffer[0:])
	if err != nil {
		s.Peer.Close()
		return nil, errors.New("error getting external response, close at " + s.Peer.RemoteAddr().String() + ", error: " + err.Error())
	}
	if n == 0 {
		return nil, errors.New("error getting external response")
	}

	return buffer, nil
}
