package timeservice

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const timeoutSeconds = 5

type Server struct {
	host    string
	log     *log.Logger
	timeout time.Duration
}

func NewServer(port string) *Server {
	port = fmt.Sprintf(":%s", port)
	return &Server{
		host:    port,
		log:     log.New(os.Stdout, "[time_server] ", log.Ldate|log.Ltime),
		timeout: time.Second * timeoutSeconds,
	}
}

func (s *Server) Run() error {
	s.log.Printf("starting server... %s", s.host)
	listener, err := net.Listen("tcp", s.host)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.log.Printf("get error on accept connection: %s", err.Error())
			continue
		}
		s.log.Printf("get connection from: %s", conn.RemoteAddr().String())
		conn.SetReadDeadline(time.Now().Add(s.timeout))
		go s.handler(conn)
	}

}

func (s *Server) handler(conn net.Conn) {
	_, err := conn.Write(encodeTime(time.Now()))
	if err != nil {
		s.log.Printf("conn write error: %s", err.Error())
		return
	}
	defer func() {
		s.log.Printf("closing connection from: %s", conn.RemoteAddr().String())
		conn.Close()
	}()
}
