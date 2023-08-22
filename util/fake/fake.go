package fake

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func CreateHTTPClient(roundTripper func(*http.Request) (*http.Response, error)) *http.Client {
	return &http.Client{
		Transport: roundTripperFunc(roundTripper),
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func SuccessJsonBody(data interface{}) io.ReadCloser {
	respWrapper := struct {
		Code int32
		Msg  string
		Data interface{}
	}{
		Code: 0,
		Msg:  "success",
		Data: data,
	}

	responseBody, err := json.Marshal(respWrapper)
	if err != nil {
		panic(err)
	}
	return io.NopCloser(bytes.NewReader(responseBody))
}
