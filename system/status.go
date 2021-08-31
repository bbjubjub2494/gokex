package system

import (
	"encoding/json"
	"github.com/lourkeur/gokex/rest"
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

func Status(rest rest.Handle) ([]StatusData, error) {
	env, err := rest.Get("system/status")
	if err != nil {
		return nil, err
	}

	dst := make([]StatusData, len(env.Data))
	for i := range env.Data {
		err = json.Unmarshal(env.Data[i], &dst[i])
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}
