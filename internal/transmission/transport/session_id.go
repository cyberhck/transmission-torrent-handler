package transport

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type SessionID struct {
	InnerTransport http.RoundTripper
	latestSessionID string
}

func (s *SessionID) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("X-Transmission-Session-Id", s.latestSessionID)
	request.Header.Set("X-Requested-With", "torrent-handler")
	req := cloneRequest(request)
	response, err := s.InnerTransport.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 409 {
		return response, nil
	}
	// get the session id from the response and continue request
	sessionID := response.Header.Get("X-Transmission-Session-Id")
	req.Header.Set("X-Transmission-Session-Id", sessionID)
	// save the latest session id
	s.latestSessionID = sessionID
	return s.InnerTransport.RoundTrip(req)
}

func cloneBody(body io.ReadCloser) (io.ReadCloser, io.ReadCloser) {
	if body == nil {
		return nil, nil
	}
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, nil
	}
	originalBody := ioutil.NopCloser(bytes.NewBuffer(buf))
	clonedBody := ioutil.NopCloser(bytes.NewBuffer(buf))

	return clonedBody, originalBody
}

func cloneRequest(r *http.Request) *http.Request {
	cloned, original := cloneBody(r.Body)
	r.Body = original
	clonedRequest := r.Clone(r.Context())
	clonedRequest.Body = cloned

	return clonedRequest
}
