package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const endpoint = "https://meta-dev.heeser.io/v1/"

type Client struct {
	httpClient *http.Client
	endpoint   string
	apiKey     string

	Auth    AuthClient
	Event   EventClient
	Project ProjectClient
	User    UserClient
}

type AddHeaderTransport struct {
	T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "go")
	return adt.T.RoundTrip(req)
}

func NewAddHeaderTransport(T http.RoundTripper) *AddHeaderTransport {
	if T == nil {
		T = http.DefaultTransport
	}
	return &AddHeaderTransport{T}
}

func NewClient() Client {
	c := Client{
		endpoint: endpoint,
	}
	return c
}

func WithAPIKey(apiKey string) *Client {
	c := &Client{
		endpoint:   endpoint,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}

	c.Auth = AuthClient{
		client: c,
	}
	c.Event = EventClient{
		client: c,
	}
	c.Project = ProjectClient{
		client: c,
	}
	c.User = UserClient{
		client: c,
	}
	return c
}

func (c *Client) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	url := c.endpoint + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	// req.Header.Set("User-Agent", c.userAgent)
	if c.apiKey != "" {
		req.Header.Set("Authorization", c.apiKey)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req = req.WithContext(ctx)
	return req, nil
}

// Do performs an HTTP request against the API.
func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	// var retries int
	var body []byte
	var err error
	if r.ContentLength > 0 {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			r.Body.Close()
			return nil, err
		}
		r.Body.Close()
	}
	for {
		if r.ContentLength > 0 {
			r.Body = ioutil.NopCloser(bytes.NewReader(body))
		}

		resp, err := c.httpClient.Do(r)
		if err != nil {
			return nil, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			return resp, err
		}
		resp.Body.Close()
		resp.Body = ioutil.NopCloser(bytes.NewReader(body))

		if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
			err = ErrorFromBody(body)
			return resp, err
		}
		if v != nil {
			if w, ok := v.(io.Writer); ok {
				_, err = io.Copy(w, bytes.NewReader(body))
			} else {
				err = json.Unmarshal(body, v)
			}
		}

		return resp, err
	}
}
