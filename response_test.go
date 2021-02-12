package main

import (
	"testing"
)

func TestResponse_Output(t *testing.T) {
	testCases := []struct {
		name           string
		resp           *Response
		expectedString string
	}{
		{
			name: "Test Ok response output",
			resp: &Response{
				SvcID:      2,
				BodyLen:    18,
				ReqID:      111,
				ReturnCode: 0,
				ClientID:   "idy",
				ClientType: 11,
				Username:   "usk",
				ExpiresIn:  1880,
				UserID:     101,
			},
			expectedString: "client_id: idy\nclient_type: 11\n" +
				"expires_in: 1880\nuser_id: 101\nusername: usk",
		},
		{
			name: "Test Ok response with error",
			resp: &Response{
				SvcID:      2,
				BodyLen:    18,
				ReqID:      111,
				ReturnCode: 2,
				Err: "db error",
			},
			expectedString: "error: CUBE_OAUTH2_ERR_DB_ERROR\nmessage: db error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.resp.Output()
			if str != tc.expectedString {
				t.Errorf("Response output Err\ngot :\n%s\nwant:\n%s", str, tc.expectedString)
			}
			return
		})
	}
}
