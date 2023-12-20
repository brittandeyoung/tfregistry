package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/module"
)

func (c *Client) CreateModule(item module.Module) (*module.Module, error) {
	rb, err := json.Marshal(item)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/modules", c.HostURL), bytes.NewBuffer(rb))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	module := new(module.Module)
	err = json.Unmarshal(body, &module)

	if err != nil {
		return nil, err
	}

	return module, nil
}