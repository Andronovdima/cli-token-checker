package main

import (
	"errors"
	"strconv"
)

type Client struct {
	Host  string
	Port  int32
	Token string
	Scope string
	ReqID int32
	Conn  Connection
}

func (c *Client) ParseArgs(args []string) error {
	port, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		return errors.New("wrong format of port")
	}
	c.Host = args[0]
	c.Port = int32(port)
	c.Token = args[2]
	c.Scope = args[3]

	return nil
}

func (c *Client) Request() (*Response, error) {
	addr := c.Host + ":" + strconv.Itoa(int(c.Port))

	err := c.Conn.Connect(addr)
	if err != nil {
		return nil, errors.New("not connect, timeout")
	}

	err = c.Conn.Write(c)
	if err != nil {
		return nil, errors.New("error with write to socket")
	}

	resp, err := c.Conn.Read()
	if err != nil {
		return nil, errors.New("error with read from socket")
	}

	err = c.Conn.Close()

	return resp, err
}
