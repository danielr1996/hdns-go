package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithToken(t *testing.T) {
	token := "mytoken"
	c := New().WithToken(token)
	assert.Equal(t, token, c.token)
}

func TestWithEndpoint(t *testing.T) {
	endpoint := "https://localhost:8080/api/v1"
	c := New().WithEndpoint(endpoint)
	assert.Equal(t, endpoint, c.endpoint)
}
func TestBaseAPI(t *testing.T) {
	host := "localhost:8080"
	endpoint := fmt.Sprintf("https://%s/api/v1", host)
	token := "mytoken"
	c := New().WithEndpoint(endpoint).WithToken(token)
	req, _ := c.baseApi().Request()
	assert.Equal(t, host, req.Host)
	assert.Equal(t, token, req.Header.Get("AUTH-API-TOKEN"))
}
