package main

import (
	"bytes"
	"testing"
)

func TestConnection_StringWrite(t *testing.T) {
	var buf []byte
	testCases := []struct {
		w             *bytes.Buffer
		name          string
		str           string
		expectedError error
		expectedBytes []byte
	}{
		{
			w:             bytes.NewBuffer(buf),
			name:          "Test Ok String write",
			str:           "testy",
			expectedError: nil,
			expectedBytes: []byte{05, 00, 00, 00, 116, 101, 115, 116, 121},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := WriteString(tc.w, tc.str)
			if tc.expectedError == nil {
				if !bytes.Equal(tc.expectedBytes, tc.w.Bytes()) {
					t.Errorf("Response output Err\ngot :\n%v\nwant:\n%v", tc.expectedBytes, tc.w.Bytes())
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

func TestConnection_StringRead(t *testing.T) {
	buf := []byte{116, 101, 115, 116, 121}
	testCases := []struct {
		w              *bytes.Buffer
		name           string
		expectedError  error
		expectedString string
	}{
		{
			w:              bytes.NewBuffer(buf),
			name:           "Test Ok String read",
			expectedError:  nil,
			expectedString: "testy",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str, err := ReadString(tc.w, int32(len(tc.expectedString)))
			if tc.expectedError == nil {
				if tc.expectedString != str {
					t.Errorf("Response output Err\ngot :\n%v\nwant:\n%v", err, tc.expectedError)
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
