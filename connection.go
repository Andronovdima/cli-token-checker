package main

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

const (
	SvcID    int32 = 0x00000002
	SvcMsg32 int32 = 0x00000001
	timeout        = 2 * time.Second
	intSize        = 4
)

type Connection interface {
	Connect(address string) error
	Write(client *Client) error
	Read() (*Response, error)
	Close() error
}

type TCPConn struct {
	conn net.Conn
}

func (c *TCPConn) Connect(address string) error {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *TCPConn) Write(client *Client) error {
	lenReq := int32(len(client.Token) + len(client.Scope) + 3*intSize)

	err := binary.Write(c.conn, binary.LittleEndian, SvcID)
	if err != nil {
		return err
	}

	err = binary.Write(c.conn, binary.LittleEndian, lenReq)
	if err != nil {
		return err
	}

	err = binary.Write(c.conn, binary.LittleEndian, client.ReqID)
	if err != nil {
		return err
	}

	err = binary.Write(c.conn, binary.LittleEndian, SvcMsg32)
	if err != nil {
		return err
	}

	err = WriteString(c.conn, client.Token)
	if err != nil {
		return err
	}

	err = WriteString(c.conn, client.Scope)
	if err != nil {
		return err
	}

	return err
}

func (c *TCPConn) Read() (*Response, error) {
	resp := Response{}
	var lenStr int32

	err := binary.Read(c.conn, binary.LittleEndian, &resp.SvcID)
	if err != nil {
		return nil, err
	}

	err = binary.Read(c.conn, binary.LittleEndian, &resp.BodyLen)
	if err != nil {
		return nil, err
	}

	err = binary.Read(c.conn, binary.LittleEndian, &resp.ReqID)
	if err != nil {
		return nil, err
	}

	err = binary.Read(c.conn, binary.LittleEndian, &resp.ReturnCode)
	if err != nil {
		return nil, err
	}

	if resp.ReturnCode == 0 {
		err = binary.Read(c.conn, binary.LittleEndian, &lenStr)
		if err != nil {
			return nil, err
		}

		resp.ClientID, err = ReadString(c.conn, lenStr)
		if err != nil {
			return nil, err
		}

		err = binary.Read(c.conn, binary.LittleEndian, &resp.ClientType)
		if err != nil {
			return nil, err
		}

		err = binary.Read(c.conn, binary.LittleEndian, &lenStr)
		if err != nil {
			return nil, err
		}

		resp.Username, err = ReadString(c.conn, lenStr)
		if err != nil {
			return nil, err
		}

		err = binary.Read(c.conn, binary.LittleEndian, &resp.ExpiresIn)
		if err != nil {
			return nil, err
		}

		err = binary.Read(c.conn, binary.LittleEndian, &resp.UserID)
		if err != nil {
			return nil, err
		}

	} else {
		err = binary.Read(c.conn, binary.LittleEndian, &lenStr)
		if err != nil {
			return nil, err
		}

		resp.Err, err = ReadString(c.conn, lenStr)
		if err != nil {
			return nil, err
		}
	}

	return &resp, nil
}

func (c *TCPConn) Close() error {
	return c.conn.Close()
}

func WriteString(w io.Writer, str string) error {
	err := binary.Write(w, binary.LittleEndian, int32(len(str)))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, []byte(str))
	return err
}

func ReadString(w io.Reader, strLen int32) (string, error) {
	var bytes []byte
	var b byte
	for i := 0; i < int(strLen); i++ {
		err := binary.Read(w, binary.LittleEndian, &b)
		if err != nil {
			return "", err
		}

		bytes = append(bytes, b)
	}

	return string(bytes), nil
}
