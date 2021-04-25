package hdns

import (
	"github.com/dghubble/sling"
	"net/http"
)

const Endpoint = "https://dns.hetzner.com/api/v1/"

type Client struct {
	token      string
	endpoint   string
	httpClient http.Client
}

func NewClient() *Client {
	c := new(Client)
	c.httpClient = http.Client{}
	c.endpoint = Endpoint
	return c
}

func (c *Client) WithToken(token string) *Client {
	c.token = token
	return c
}

func (c *Client) WithEndpoint(endpoint string) *Client {
	c.endpoint = endpoint
	return c
}

func (c *Client) BaseApi() *sling.Sling {
	return sling.
		New().
		Base(c.endpoint).
		Set("Auth-API-Token", c.token)
}
