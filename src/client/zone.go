package client

import (
	"errors"
	"fmt"
	"github.com/danielr1996/hdns-go/schema"
)

// ZoneClient provides methods to query zones in the Hetzner DNS API
type ZoneClient struct {
	Client *Client
}

// GetAll receives all zones associated with the user from the Hetzner DNS API
func (c *ZoneClient) GetAll() ([]schema.Zone, error) {
	success := new(schema.ZoneResponse)
	res, err := c.Client.baseApi().Get("zones").ReceiveSuccess(success)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return success.Zones, nil
}

func (c *ZoneClient) getAllByName(name string) ([]schema.Zone, error) {
	success := new(schema.ZoneResponse)
	failure := new(schema.ErrorResponse)
	resp, err := c.Client.baseApi().Get("zones").QueryStruct(&schema.ZoneParams{Name: name}).Receive(success, failure)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 399){
		return nil, errors.New(failure.Error.Message)
	}
	return success.Zones, err
}

// GetByName receives the zone with the specified name from the Hetzner DNS API
func (c *ZoneClient) GetByName(name string) (schema.Zone, error) {
	zones, err := c.getAllByName(name)
	if err != nil {
		return schema.Zone{}, err
	}
	return zones[0], nil
}
