package api_test

import (
	"errors"
	"io"
)

func mockReadError(reader io.Reader) ([]byte, error) {
	return nil, errors.New("mock read error")
}

func mockMarshalError(v interface{}) ([]byte, error) {
	return nil, errors.New("mock marshal error")
}
