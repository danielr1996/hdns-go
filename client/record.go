package client

import (
	"errors"
	"fmt"
	"github.com/danielr1996/hdns-go/schema"
	"strings"
)

// RecordClient provides methods to query and modify records in the Hetzner DNS API
type RecordClient struct {
	Client *Client
}

// GetAll receives all records associated with the user from the Hetzner DNS API
func (c *RecordClient) GetAll() ([]schema.Record, error) {
	success := new(schema.RecordResponse)
	resp, err := c.Client.baseApi().Get("records").ReceiveSuccess(success)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 399) {
		return nil, errors.New(resp.Status)
	}
	return success.Records, err
}

// GetById receives the record with the specified id from the Hetzner DNS API
func (c *RecordClient) GetById(id string) (schema.Record, error) {
	success := new(schema.RecordResponse)
	failure := new(schema.ErrorResponse)
	resp, err := c.Client.baseApi().Get("records/").Path(id).Receive(success, failure)
	if err != nil {
		return schema.Record{}, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 399) {
		return schema.Record{}, errors.New(failure.Error.Message)
	}
	return success.Record, err
}

// Create created a new Record in the Hetzner DNS API.
// Allowed recordTypes are A, AAA, NS, MX, CNAME, RP, TXT, SOA, HINFO, SRV, DANE, TLSA, DS, CAA
func (c *RecordClient) Create(name string, recordType string, value string, zoneId string) (schema.Record, error) {
	maxRetry := 10
	body := &schema.Record{
		Name:   name,
		Type:   recordType,
		Value:  value,
		ZoneId: zoneId,
	}
	success := new(schema.RecordResponse)
	failure := new(schema.ErrorResponse)
	resp, err := c.Client.baseApi().Post("records").BodyJSON(body).Receive(success, failure)

	if err != nil {
		return schema.Record{}, err
	}

	// Hetzner DNS API returns HTTP Status 422 with message "409 Conflict: " when creating multiple records in parallel
	// (e.g. in a goroutine) for every record but the first created. It seems that the second call shortly afterwards works,
	// so we just retry `maxRetry` times.
	retries := 0
	for resp.StatusCode == 422 && strings.Contains(failure.Error.Message, "409 Conflict") && retries < maxRetry {
		retries++
		resp, err = c.Client.baseApi().Post("records").BodyJSON(body).Receive(success, failure)
	}
	if err != nil {
		return schema.Record{}, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 399) {
		return schema.Record{}, errors.New(fmt.Sprintf("Code: %i; Message: %s", failure.Error.Code, failure.Error.Message))
	}
	return success.Record, nil
}

// Update updates a record in the Hetzner DNS API
func (c *RecordClient) Update(name string, recordType string, value string, zoneId string, id string) (schema.Record, error) {
	maxRetry := 10
	body := &schema.Record{
		Name:   name,
		Type:   recordType,
		Value:  value,
		ZoneId: zoneId,
	}
	success := new(schema.RecordResponse)
	failure := new(schema.ErrorResponse)
	resp, err := c.Client.baseApi().Put("records/").Path(id).BodyJSON(body).Receive(success, failure)

	if err != nil {
		return schema.Record{}, err
	}

	// Hetzner DNS API returns HTTP Status 422 with message "409 Conflict: " when creating multiple records in parallel
	// (e.g. in a goroutine) for every record but the first created. It seems that the second call shortly afterwards works,
	// so we just retry `maxRetry` times.
	retries := 0
	for resp.StatusCode == 422 && strings.Contains(failure.Error.Message, "409 Conflict") && retries < maxRetry {
		retries++
		resp, err = c.Client.baseApi().Put("records/").Path(id).BodyJSON(body).Receive(success, failure)
	}
	if err != nil {
		return schema.Record{}, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 399) {
		return schema.Record{}, errors.New(fmt.Sprintf("Code: %i; Message: %s", failure.Error.Code, failure.Error.Message))
	}
	return success.Record, nil
}

// Delete deletes a record in the Hetzner DNS API
func (c *RecordClient) Delete(id string) error {
	resp, err := c.Client.baseApi().Delete("records/").Path(id).ReceiveSuccess(nil)
	if err != nil {
		return err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 399) {
		return errors.New(resp.Status)
	}
	return nil
}
