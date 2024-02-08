package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	methodGetUpdates  = "getUpdates"
	methodSendMessage = "sendMessage"
)

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: NewBasePath(token),
		client:   http.Client{},
	}
}

func NewBasePath(token string) string {
	return "bot" + token
}

// send message

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Set("offset", strconv.Itoa(offset))
	q.Set("limit", strconv.Itoa(limit))

	data, err := c.doRequest(methodGetUpdates, q)
	if err != nil {
		return nil, err
	}

	var res UpdateResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("can't decode Update: %w", err)
	}
	return res.Result, nil
}

func (c *Client) SendMessage(chatId int, text string) error {
	q := url.Values{}
	q.Set("chat_id", strconv.Itoa(chatId))
	q.Set("text", text)

	_, err := c.doRequest(methodSendMessage, q)
	if err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}

func (c *Client) doRequest(method string, q url.Values) ([]byte, error) {
	url := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("can't generate request: %w", err)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can'r read body response: %w", err)
	}
	return data, nil
}
