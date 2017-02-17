package server

import (
	"log"
	"net"
	"strconv"
)

// 枚举值 服务器状态的枚举值
const (
	Close = iota
	Open
)

// Server  服务器的结构
type Server struct {
	port     int            // 服务器端口
	status   int            // 服务器状态（枚举值）
	listener net.Listener   // 侦听器
	handler  func(net.Conn) // 处理链接的函数（作为Start()的参数传入）
}

// Init 设置端口，并返回Server指针。并不会启动服务器。
func Init(port int) *Server {
	s := new(Server)
	s.port = port
	return s
}

// GetStatus 返回服务器状态（枚举值）
func (s *Server) GetStatus() int {
	return s.status
}

// Start server
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
	s.status = Close
	s.listener.Close()
}
