package hdns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

type ZoneResponse struct {
	Zones []map[string]interface{} `json:"zones"`
}

type RecordsResponse struct {
	Records []map[string]interface{} `json:"records"`
}
type RecordResponse struct {
	Record map[string]string `json:"record"`
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
	var res ZoneResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return res.Zones
}

func (c Client) ZonesByName(name string) []map[string]interface{} {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/zones?name=%s", c.endpoint, name), nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
	req.Header.Add("Auth-API-Token", c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	var res ZoneResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return res.Zones
}

func (c Client) Records() []map[string]interface{} {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/records", c.endpoint), nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
	req.Header.Add("Auth-API-Token", c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	var res RecordsResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return res.Records
}

//recordType A, AAA, NS, MX, CNAME, RP, TXT, SOA, HINFO, SRV, DANE, TLSA, DS, CAA
func (c Client) AddRecord(name string, recordType string, value string, zone_id string) map[string]string {
	values := map[string]string{"name": name, "type": recordType, "value": value, "zone_id": zone_id}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/records", c.endpoint), bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Printf(err.Error())
	}
	req.Header.Add("Auth-API-Token", c.token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	var res RecordResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return res.Record
}
