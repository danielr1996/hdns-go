package hdns

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL string = "https://dns.hetzner.com/api/v1"

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

type Zone struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ZonesResponse struct {
	Zones []Zone `json:"zones"`
}

type ZoneResponse struct {
	Zone Zone `json:"zone"`
}

type Record struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RecordsResponse struct {
	Records []Record `json:"records"`
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Auth-API-Token", s.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (s *Client) GetZones() (*[]Zone, error) {
	URL := fmt.Sprintf("%s/zones", baseURL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data ZonesResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data.Zones, nil
}

func (s *Client) GetZone(zoneid string) (*Zone, error) {
	URL := fmt.Sprintf("%s/zones/%s", baseURL, zoneid)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data ZoneResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data.Zone, nil
}

func (s *Client) GetZoneByName(name string) (*Zone, error) {
	URL := fmt.Sprintf("%s/zones?name=%s", baseURL, name)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data ZonesResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data.Zones[0], nil
}

func (s *Client) GetRecords(zoneid string) (*[]Record, error) {
	URL := fmt.Sprintf("%s/records?zone_id=%s", baseURL, zoneid)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data RecordsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data.Records, nil
}
