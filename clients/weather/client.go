package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	host     string
	basePath string
	apiKey   string
	client   http.Client
}

func New(host, basePath string, apiKey string) *Client {
	return &Client{
		host:     host,
		basePath: basePath,
		apiKey:   apiKey,
		client:   http.Client{},
	}
}

func (c *Client) doRequest(query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   c.basePath,
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("can't create request: %w", err)
	}

	req.URL.RawQuery = query.Encode()
	// log.Print(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %w", err)
	}

	return res, nil
}
