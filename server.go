package timeservice

import (
	"fmt"
	"log"
	"os"
)

type Server struct {
	host string
	log  *log.Logger
}

func NewServer(port string) *Server {
	port = fmt.Sprintf(":%s", port)
	return &Server{
		host: port,
		log:  log.New(os.Stdout, "TimeServer", log.LUTC),
	}
}

func (s *Server) Run() {

}

func (s *Server) handler() {

}
