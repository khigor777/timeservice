package timeservice

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const (
	timeoutConnectionsSeconds = 5
)

type Server struct {
	sync.Mutex
	host        string
	log         *log.Logger
	connections map[net.Conn]struct{}
}

func NewServer(port string) *Server {
	port = fmt.Sprintf(":%s", port)
	return &Server{
		host: port,
		log:  log.New(os.Stdout, "[time_server] ", log.Ldate|log.Ltime),
	}
}

//run server
func (s *Server) Run() {

	s.log.Printf("starting server... %s", s.host)
	listener, err := net.Listen("tcp", s.host)
	if err != nil {
		s.log.Fatalf("start server error: %s", err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.log.Printf("get error on accept connection: %s", err.Error())
			continue
		}
		s.log.Printf("get connection from: %s", conn.RemoteAddr().String())

		conn.SetReadDeadline(time.Now().Add(timeoutConnectionsSeconds))
		s.addConn(conn)
		go s.handler(conn)
	}
}

//handle connection
func (s *Server) handler(conn net.Conn) {
	_, err := conn.Write(encodeTime(time.Now()))
	if err != nil {
		s.log.Printf("conn write error: %s", err.Error())
		return
	}
	defer func() {
		s.log.Printf("closing connection from: %s", conn.RemoteAddr().String())
		s.delConn(conn)
		conn.Close()
	}()
}

//close all connections
func (s *Server) Shutdown(ctx context.Context) {
	s.log.Printf("shutdown ... ")
	for {
		select {
		case <-ctx.Done():
			s.log.Printf("closed by timeout, opened connections: %d", len(s.connections))
			return
		default:
			for k := range s.connections {
				if len(s.connections) == 0 {
					s.log.Printf("server has closed success ...")
					return
				}
				if k != nil {
					k.Close()
				}
			}
		}
	}
}

//track connection for showdown
func (s *Server) addConn(conn net.Conn) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	if s.connections == nil {
		s.connections = make(map[net.Conn]struct{})
	}
	s.connections[conn] = struct{}{}
}

//del connection after handle it
func (s *Server) delConn(conn net.Conn) {
	if _, ok := s.connections[conn]; ok {
		delete(s.connections, conn)
	}
}
