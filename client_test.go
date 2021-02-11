package main

import (
	"errors"
	"testing"
)

func TestClient_ParseArgs(t *testing.T) {
	client := Client{}
	testCases := []struct {
		name           string
		arguments      []string
		expectedError  error
		expectedClient Client
	}{
		{
			name:          "Right format of args",
			arguments:     []string{"cube.testserver.mail.ru", "4995", "abracadabra", "test"},
			expectedError: nil,
			expectedClient: Client{
				Host:  "cube.testserver.mail.ru",
				Port:  4995,
				Token: "abracadabra",
				Scope: "test",
			},
		},
		{
			name:           "Bad format of Port",
			arguments:      []string{"cube.testserver.mail.ru", "err_port", "abracadabra", "test"},
			expectedError:  errors.New("wrong format of port"),
			expectedClient: Client{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := client.ParseArgs(tc.arguments)
			if tc.expectedError == nil {
				if tc.expectedClient != client {
					t.Errorf("Response output Err\ngot :\n%v\nwant:\n%v", client, tc.expectedClient)
				}
				return
			} else {
				if tc.expectedError.Error() != err.Error() {
					t.Errorf("Response output Err\ngot :\n%v\nwant:\n%v", err, tc.expectedError)
				}
				return
			}

			t.Fatal()
		})
	}

}

func TestClient_Request(t *testing.T) {
	testCases := []struct {
		name             string
		client           Client
		expectedError    error
		expectedResponse Response
	}{
		{
			name: "Connect Error",
			client: Client{
				Host: "blblblbb",
				Port: 1500,
				Conn: &ConnMock{},
			},
			expectedError:    errors.New("not connect, timeout"),
			expectedResponse: Response{},
		},
		{
			name: "Test Ok",
			client: Client{
				Host:  "cube.testserver.mail.ru",
				Port:  4995,
				Token: "abracadabra",
				Scope: "test",
				ReqID: 323,
				Conn:  NewTestOkConnMock(),
			},
			expectedError: nil,
			expectedResponse: Response{
				SvcID:      2,
				BodyLen:    58,
				ReqID:      323,
				ReturnCode: 0,
				ClientID:   "test_client_id",
				ClientType: 20100,
				Username:   "testuser@mail.ru",
				ExpiresIn:  3600,
				UserID:     101010,
				Err:        "",
			},
		},
		{
			name: "Test Token Not Found ",
			client: Client{
				Host:  "cube.testserver.mail.ru",
				Port:  4995,
				Token: "abbba",
				Scope: "test",
				ReqID: 323,
				Conn:  NewTestNotFoundConnMock(),
			},
			expectedError: nil,
			expectedResponse: Response{
				SvcID:      2,
				BodyLen:    58,
				ReqID:      323,
				ReturnCode: 1,
				Err:        "token not found",
			},
		},
		{
			name: "Test Write Error",
			client: Client{
				Host:  "cube.testserver.mail.ru",
				Port:  4995,
				Token: "abbba",
				Scope: "test",
				ReqID: 323,
				Conn:  NewTestWriteErrConnMock(),
			},
			expectedError:    errors.New("error with write to socket"),
			expectedResponse: Response{},
		},
		{
			name: "Test Read Error",
			client: Client{
				Host:  "cube.testserver.mail.ru",
				Port:  4995,
				Token: "abbba",
				Scope: "test",
				ReqID: 323,
				Conn:  NewTestReadErrConnMock(),
			},
			expectedError:    errors.New("error with read from socket"),
			expectedResponse: Response{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := tc.client.Request()
			if tc.expectedError == nil {
				if tc.expectedResponse != *resp {
					t.Errorf("Response output Err\ngot :\n%v\nwant:\n%v", *resp, tc.expectedResponse)
				}
				return
			} else {
				if tc.expectedError.Error() != err.Error() {
					t.Errorf("Response output Err\ngot :\n%v\nwant:\n%v", err, tc.expectedError)
				}
				return
			}

			t.Fatal()
		})
	}

}
