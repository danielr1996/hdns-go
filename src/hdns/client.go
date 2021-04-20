package hdns

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const Endpoint = "https://dns.hetzner.com/api/v1"

type Client struct {
	token      string
	endpoint   string
	httpClient http.Client
}

func NewClient() *Client {
	c := new(Client)
	c.httpClient = http.Client{}
	c.endpoint = "https://dns.hetzner.com/api/v1"
	return c
}

func (c Client) WithToken(token string) Client {
	c.token = token
	return c
}

func (c Client) WithEndpoint(endpoint string) Client {
	c.endpoint = endpoint
	return c
}

type Response struct {
	Zones []map[string]interface{} `json:"zones"`
}

func (c Client) Zones() []map[string]interface{} {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/zones", c.endpoint), nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
	req.Header.Add("Auth-API-Token", c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	var res Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return res.Zones
}

func (c Client) AddRecord() []map[string]interface{} {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/zones", c.endpoint), nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
	req.Header.Add("Auth-API-Token", c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	var res Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return res.Zones
}
