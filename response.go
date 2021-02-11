package main

import (
	"fmt"
)

type Response struct {
	SvcID      int32
	BodyLen    int32
	ReqID      int32
	ReturnCode int32
	ClientID   string
	ClientType int32
	Username   string
	ExpiresIn  int32
	UserID     int64
	Err        string
}

func (r *Response) Output() string {
	if 0 == r.ReturnCode {
		return fmt.Sprintf("client_id: %s\n"+"client_type: %d\n"+
			"expires_in: %d\n"+"user_id: %d\n"+"username: %s",
			r.ClientID, r.ClientType, r.ExpiresIn, r.UserID, r.Username)
	} else {
		var codeErr string

		switch r.ReturnCode {
		case 1:
			codeErr = "CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND"
		case 2:
			codeErr = "CUBE_OAUTH2_ERR_DB_ERROR"
		case 3:
			codeErr = "CUBE_OAUTH2_ERR_UNKNOWN_MSG"
		case 4:
			codeErr = "CUBE_OAUTH2_ERR_BAD_PACKET"
		case 5:
			codeErr = "CUBE_OAUTH2_ERR_BAD_CLIENT"
		case 6:
			codeErr = "CUBE_OAUTH2_ERR_BAD_SCOPE"
		}

		return fmt.Sprintf("error: %s\n"+"message: %s", codeErr, r.Err)
	}
}
