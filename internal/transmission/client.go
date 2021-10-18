package transmission

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cyberhck/torrent-handler/internal/transmission/transport"
	"net/http"
	"time"
)

type Client struct {
	endpoint   string
	httpClient http.Client
}

func New(endpoint string) *Client {
	httpClient := http.Client{
		Timeout:   time.Second * 3,
		Transport: &transport.SessionID{InnerTransport: http.DefaultTransport},
	}
	return &Client{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
}

type Arguments struct {
	Filename string `json:"filename"`
}

type AddRequest struct {
	Method    string    `json:"method"`
	Arguments Arguments `json:"arguments"`
}

func (c Client) AddTorrent(url string) error {
	path := fmt.Sprintf("%s/transmission/rpc", c.endpoint)
	buffer := &bytes.Buffer{}
	err := json.NewEncoder(buffer).Encode(&AddRequest{
		Method:    "torrent-add",
		Arguments: Arguments{Filename: url},
	})
	if err != nil {
		return err
	}
	res, err := c.httpClient.Post(path, "application/json", buffer)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("error status: %d", res.StatusCode)
	}
	return nil
}
