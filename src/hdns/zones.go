package hdns

import (
	"errors"
)

type Zone struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Ttl         int      `json:"ttl"`
	Registrar   string   `json:"registrar"`
	Ns          []string `json:"ns"`
	Created     string   `json:"created"`
	Verified    string   `json:"verified"`
	Modified    string   `json:"modified"`
	Project     string   `json:"project"`
	Owner       string   `json:"owner"`
	Permissions string   `json:"permissions"`
	ZoneType    struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Prices      string `json:"prices"`
	} `json:"zone_type"`
	Status          string `json:"status"`
	Paused          bool   `json:"paused"`
	IsSecondaryDns  bool   `json:"is_secondary_dns"`
	TxtVerification struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	} `json:"txt_verification"`
	RecordsCount int `json:"records_count"`
}

type Meta struct {
	Pagination struct {
		Page         int `json:"page"`
		PerPage      int `json:"per_page"`
		PreviousPage int `json:"previous_page"`
		NextPage     int `json:"next_page"`
		LastPage     int `json:"last_page"`
		TotalEntries int `json:"total_entries"`
	} `json:"pagination"`
}

type ZoneSuccess struct {
	Zones []Zone `json:"zones"`
	Meta  Meta   `json:"meta"`
}

type ZonesError struct {
	Meta  Meta   `json:"meta"`
	Error struct {
		Message string`json:"message"`
		Code int`json:"code"`
	}`json:"error"`
}

type ZonesParams struct {
	Name       string `url:"name,omitempty"`
	Page       string `url:"page,omitempty"`
	PerPage    string `url:"per_page,omitempty"`
	SearchName string `url:"search_name,omitempty"`
}

func (c *Client) Zones() ([]Zone, error) {
	zones := new(ZoneSuccess)
	_, err := c.BaseApi().Get("zones").ReceiveSuccess(zones)
	if err != nil {
		return nil, err
	}
	return zones.Zones, nil
}

func (c *Client) ZonesByName(name string) ([]Zone, error) {
	success := new(ZoneSuccess)
	failure := new(ZonesError)
	resp, err := c.BaseApi().Get("zones").QueryStruct(&ZonesParams{Name: name}).Receive(success, failure)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 399){
		return nil, errors.New(failure.Error.Message)
	}
	return success.Zones, err
}

func (c *Client) ZoneByName(name string) (Zone, error) {
	zones, err := c.ZonesByName(name)
	if err != nil {
		return Zone{}, err
	}
	return zones[0], nil
}
