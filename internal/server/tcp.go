package server

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/AafilUmar/gocha/internal/store"
)

type Server struct {
	addr  string
	cache *store.Cache
}


func New(arr string, s *store.Cache) *Server {
	return &Server{
		addr:  arr,
		cache: s,
	}
}

func (s *Server) Start() error {
	listner, err := net.Listen("tcp", s.addr)

	if err != nil {
		return err
	}
	defer listner.Close()

	log.Println("Server listening on", s.addr)

	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}

		go s.handleConnection(conn)

	}

}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			return
		}
		line = strings.TrimSpace(line)
		response := s.handleCommand(line)
		conn.Write([]byte(response + "\n"))

	}
}

func (s *Server) handleCommand(line string) string {
	parts := strings.Split(line," ")

	if len(parts) == 0 {
	return "ERR empty command"
	}
	switch strings.ToUpper(parts[0]){

		case "SET":
		if len(parts) < 3 {
			return "ERR usage : SET key value"
		}
		s.cache.Set(parts[1],parts[2],0)
		return "OK"
		case "GET":
		if len(parts) < 2 {
		return "ERR usage: GET key"
		}
		val,ok := s.cache.Get(parts[1])

		if !ok {
		return "NULL"
		}
		return val

		case "DEL" :
			if len(parts) < 2 {
			return "ERR usage: DEL key"
			}
			if s.cache.Delete(parts[1]){
			return "OK"
			}
			return "NOT FOUND"
		default :
		return "ERR unknown command"
	}
}
