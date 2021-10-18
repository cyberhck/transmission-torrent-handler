package transport

import (
	"net/http"
)

type SessionID struct {
	InnerTransport http.RoundTripper
	latestSessionID string
}

func (s *SessionID) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("X-Transmission-Session-Id", s.latestSessionID)
	request.Header.Set("X-Requested-With", "torrent-handler")
	response, err := s.InnerTransport.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 409 {
		return response, nil
	}
	// get the session id from the response and continue request
	sessionID := response.Header.Get("X-Transmission-Session-Id")
	request.Header.Set("X-Transmission-Session-Id", sessionID)
	// save the latest session id
	s.latestSessionID = sessionID
	return s.InnerTransport.RoundTrip(request)
}
