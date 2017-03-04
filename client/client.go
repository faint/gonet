package client

import (
	"net"
	"time"
)

// the status of server
const (
	Lost = iota
	Stop
	Open
	Waring
	Full
)

const (
	retryDelay = 1 // time.Second, to retry to dial client.
)

// Client save the other server as client.
type Client struct {
	address string // ip +":"+ port
	status  int
	conn    net.Conn
}

// GetClient return client pointer and set the address.
// but did not dial to target.
func GetClient(ip, port string) *Client {
	c := new(Client)

	if ip == "" {
		ip = "127.0.0.1"
	}
	c.address = ip + ":" + port

	c.status = Lost

	for { // when can't connect, retry every 1 Second
		conn, err := c.Dial()
		if err == nil {
			c.conn = conn
			break
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
	c.conn = conn
	c.status = Open
	return conn, nil
}

// Send message use conn
func (c *Client) Send(b []byte) (int, error) {
	n, e := c.conn.Write(b)
	return n, e
}

// Close the client
func (c *Client) Close() {
	c.status = Stop
	c.conn.Close()
}
