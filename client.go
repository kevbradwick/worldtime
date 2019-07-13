package worldtime

import (
	"fmt"
	"net/http"
)

// ClientConfig holds values to configure the client
type ClientConfig struct {
	BaseURL string
}

// Client is the struct used to interact with the WorldTime API
type Client struct {
	baseURL string
}

// Timezones will return a list of timezones
func (c *Client) Timezones() ([]Timezone, error) {
	var tzs []Timezone
	_, err := c.sendRequest(http.MethodGet, "/timezone", &tzs)
	if err != nil {
		return nil, err
	}

	return tzs, nil
}

// Timezone will return information about the timezone such as utc offset, daylight savings and more.
func (c *Client) Timezone(area Timezone) (TimezoneInfo, error) {
	var info TimezoneInfo
	_, err := c.sendRequest(http.MethodGet, fmt.Sprintf("/timezone/%s", area), &info)
	if err != nil {
		return info, err
	}

	return info, nil
}

// NewClient will construct a new WorldTime client
func NewClient(config ClientConfig) *Client {
	return &Client{baseURL: config.BaseURL}
}

// NewDefaultClient builds a new client with default values
func NewDefaultClient() *Client {
	return NewClient(ClientConfig{BaseURL: "https://worldtimeapi.org/api"})
}
