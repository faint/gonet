package server

import "net"

// GetConn return the conn for Write.
func GetConn(ip, port string) (net.Conn, error) {
	var target string
	if ip != "" {
		target = ip + ":" + port
	}
	conn, err := net.Dial("tcp", target)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
