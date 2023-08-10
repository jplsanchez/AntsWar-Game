package sockets

import (
	"errors"
	"log"
	"net"
	"time"
)

type Message struct {
	From    string
	Payload []byte
}

type Server struct {
	ListenAddr string
	Listener   net.Listener
	Lock       chan struct{}
	Peer       net.Conn
	ReadyToGo  chan bool
}

func NewServer(listenAddr string, ready chan bool) *Server {
	return &Server{
		ListenAddr: listenAddr,
		Lock:       make(chan struct{}),
		Peer:       nil,
		ReadyToGo:  ready,
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
		s.ReadyToGo <- true
	}
}

func (s *Server) SendRequest(message []byte) ([]byte, error) {
	if s.Peer == nil {
		return nil, errors.New("no peer connected")
	}
	if _, err := s.Peer.Write(message); err != nil {
		s.Peer.Close()
		return nil, errors.New("error writing to connection" + ", error: " + err.Error())
	}

	if err := s.Peer.SetReadDeadline(time.Now().Add(15 * time.Minute)); err != nil {
		s.Peer.Close()
		return nil, errors.New("error setting read deadline" + ", error: " + err.Error())
	}

	buffer := make([]byte, 2048)
	n, err := s.Peer.Read(buffer[0:])
	if err != nil {
		s.Peer.Close()
		return nil, errors.New("error getting external response" + ", error: " + err.Error())
	}
	if n == 0 {
		return nil, errors.New("error getting external response")
	}

	return buffer, nil
}
