package utils

import (
	"bytes"
	"encoding/binary"
	"net/http"
	"net/http/httptest"
	"github.com/sirupsen/logrus"
)

func HttpTestRequest(handler http.Handler, method, url string, payload []byte) (*httptest.ResponseRecorder, *http.Request) {

	request := make(chan *http.Request, 1)
	errors := make(chan error, 1)

	if binary.Size(payload) > 0 {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
		request <- req
		errors <- err
	} else {
		req, err := http.NewRequest(method, url, nil)
		request <- req
		errors <- err
	}

	if <-errors != nil {
		logrus.Error(<-errors)
	}

	rr := httptest.NewRecorder()

	return rr, <-request
}
