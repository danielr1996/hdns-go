package client

import (
	"github.com/dghubble/sling"
	"net/http"
)

// Endpoint is the default endpoint for the Hetzner DNS API
const Endpoint = "https://dns.hetzner.com/api/v1/"

// Client provides methods to configure the connection to the Hetzner DNS API
type Client struct {
	token      string
	endpoint   string
	httpClient http.Client
	Record     *RecordClient
	Zone       *ZoneClient
}

// New instantiates a new Client for the Hetzner DNS API
func New() *Client {
	c := new(Client)
	c.httpClient = http.Client{}
	c.endpoint = Endpoint
	c.token = ""
	c.Record = &RecordClient{Client: c}
	c.Zone = &ZoneClient{Client: c}
	return c
}

// WithToken configures a Client to use the specified token
func (c *Client) WithToken(token string) *Client {
	c.token = token
	return c
}

// WithEndpoint configures a Client to use the specified endpoint
func (c *Client) WithEndpoint(endpoint string) *Client {
	c.endpoint = endpoint
	return c
}

// baseApi provides a configures rest client to be used by the resource clients
func (c *Client) baseApi() *sling.Sling {
	return sling.
		New().
		Base(c.endpoint).
		Set("Auth-API-Token", c.token)
}
