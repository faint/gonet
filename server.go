package server

import (
	"log"
	"net"
	"strconv"
)

// 枚举值 服务器状态的枚举值
const (
	Stop = iota
	Open
)

// Server  服务器的结构
type Server struct {
	port     int            // 服务器端口
	status   int            // 服务器状态（枚举值）
	listener net.Listener   // 侦听器
	handler  func(net.Conn) // 处理链接的函数
}

// GetInstance Get Server instance
// need port to listen
// won't Start the Server
func GetInstance(port int) *Server {
	s := new(Server)
	s.port = port
	s.status = Stop
	return s
}

// GetStatus Get server status
func (s *Server) GetStatus() int {
	return s.status
}

// Start server
// need handler to handle conn
func (s *Server) Start(handler func(net.Conn)) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		log.Fatal("Listen failed:", err)
	}
	s.listener = listener

	s.status = Open
	for s.status == Open {
		conn, err := listener.Accept()
		if err != nil {
			// Todo 日志链接错误
		}

		go handler(conn) // 利用传入的函数处理链接
	}
}

// Stop server
func (s *Server) Stop() {
	s.status = Stop
	s.listener.Close()
}
