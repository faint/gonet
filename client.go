package server

import (
	"net"
	"time"
)

const (
	retryDelay = 1 // time.Second, to retry to dial client.
)

// Client save the other server as client.
type Client struct {
	address string // ip +":"+ port
	status  int
	Conn    net.Conn
}

// GetClient return client pointer and set the address.
// but did not dial to target.
func GetClient(ip, port string) *Client {
	c := new(Client)

	if ip != "" {
		c.address = ip + ":" + port
	}

	c.status = Lost

	open := false
	for !open { // when can't connect, retry every 1 Second
		_, err := c.Dial()
		if err == nil {
			open = true
		}
		time.Sleep(retryDelay * time.Second)
	}

	return c
}

// Dial dial target address.
func (c *Client) Dial() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return nil, err
	}
	c.Conn = conn
	c.status = Open
	return conn, nil
}
