package worldtime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *Client) sendRequest(method, path string, i interface{}) (interface{}, error) {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}

	url := fmt.Sprintf("%s%s", c.baseURL, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "worldtime-go/1.0.0")

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
