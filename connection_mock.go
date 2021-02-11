package main

import (
	"errors"
)

type ConnMock struct {
	resp     *Response
	errWrite bool
	errRead  bool
}

func NewTestOkConnMock() *ConnMock {
	return &ConnMock{
		resp: &Response{
			2,
			58,
			323,
			0,
			"test_client_id",
			20100,
			"testuser@mail.ru",
			3600,
			101010,
			"",
		},
		errWrite: false,
		errRead:  false,
	}
}

func NewTestNotFoundConnMock() *ConnMock {
	return &ConnMock{
		resp: &Response{
			SvcID:      2,
			BodyLen:    58,
			ReqID:      323,
			ReturnCode: 1,
			Err:        "token not found",
		},
		errWrite: false,
		errRead:  false,
	}
}

func NewTestWriteErrConnMock() *ConnMock {
	return &ConnMock{
		resp:     nil,
		errWrite: true,
		errRead:  false,
	}
}

func NewTestReadErrConnMock() *ConnMock {
	return &ConnMock{
		resp:     nil,
		errWrite: false,
		errRead:  true,
	}
}

func (c *ConnMock) Connect(address string) error {
	if address != "cube.testserver.mail.ru:4995" {
		return errors.New("not connect, timeout")
	}
	return nil
}

func (c *ConnMock) Write(client *Client) error {
	if c.errWrite {
		return errors.New("err")
	}

	return nil
}

func (c *ConnMock) Read() (*Response, error) {
	if c.errRead {
		return nil, errors.New("err")
	}

	return c.resp, nil
}

func (c *ConnMock) Close() error {
	return nil
}
