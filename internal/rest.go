package internal

import (
	"encoding/json"
	"net/http"
)

// Common code for the REST API.
// https://www.okex.com/docs-v5/en/#rest-api

type Enveloppe struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []json.RawMessage `json:"data"`
}

const baseUrl = "https://www.okex.com/api/v5/"

var client = http.DefaultClient

type RestHandle struct{}

var Rest RestHandle

func (RestHandle) Get(endpoint string) (*Enveloppe, error) {
	resp, err := client.Get(baseUrl + "system/status")
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	env := &Enveloppe{}
	err = decoder.Decode(env)
	if err != nil {
		return nil, err
	}
	return env, nil
}
