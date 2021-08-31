package system

import (
	"encoding/json"
	"github.com/lourkeur/gokex/internal"
)

type StatusData struct {
	Title       string `json:"title"`
	State       string `json:"state"`
	Begin       string `json:"begin"`
	End         string `json:"end"`
	Href        string `json:"href"`
	ServiceType string `json:"serviceType"`
	System      string `json:"system"`
	ScheDesc    string `json:"scheDesc"`
}

func Status() (*StatusData, error) {
	data := &StatusData{}
	env, err := internal.Rest.Get("system/status")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(env.Data[0], data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
