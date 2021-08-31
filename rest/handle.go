package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Common code for the REST API.
// https://www.okex.com/docs-v5/en/#rest-api

type Enveloppe struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []json.RawMessage `json:"data"`
}

type Options struct {
	BaseUrl   string // defaults to https://www.okex.com/api/v5/
	Simulated bool   // whether to include the "simulated trading" header

	AccessKey  string // the API access key
	Passphrase string // the passphrase
	SecretKey  string // the secret key
}

type Handle interface {
	Get(endpoint string) (*Enveloppe, error)
	Post(enpoint string, body interface{}) (*Enveloppe, error)
}

type handle struct {
	Options
	http *http.Client
}

func NewHandle(opts *Options) (Handle, error) {
	this := &handle{*opts, &http.Client{}}
	if this.BaseUrl == "" {
		this.BaseUrl = "https://www.okex.com/api/v5/"
	}
	if this.BaseUrl[len(this.BaseUrl)-1] != '/' {
		this.BaseUrl += "/"
	}
	return this, nil
}

func (this *handle) addHeaders(req *http.Request) {
	if this.Simulated {
		req.Header.Set("x-simulated-trading", "1")
	}
	req.Header.Set("ok-access-key", this.AccessKey)
	req.Header.Set("ok-access-passphrase", this.Passphrase)
}

func (this *handle) Get(endpoint string) (*Enveloppe, error) {
	req, err := http.NewRequest("GET", this.BaseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}
	this.addHeaders(req)
	resp, err := this.http.Do(req)
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
	if env.Code != "0" {
		return nil, fmt.Errorf("OKEx error: %s - %s", env.Code, env.Msg)
	}
	return env, nil
}

func (this *handle) Post(endpoint string, body interface{}) (*Enveloppe, error) {
	bodyMsg, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", this.BaseUrl+endpoint, bytes.NewBuffer(bodyMsg))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	this.addHeaders(req)
	{
		// sign request per https://www.okex.com/docs-v5/en/#rest-api-authentication-signature
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")

		message := []byte{}
		message = append(message, timestamp...)
		message = append(message, "POST"...)
		message = append(message, ("/api/v5/" + endpoint)...)
		message = append(message, bodyMsg...)

		secret := []byte{}
		secret = append(secret, this.SecretKey...)
		hmac := hmac.New(sha256.New, secret)
		hmac.Write(message)
		req.Header.Set("ok-access-sign", base64.StdEncoding.EncodeToString(hmac.Sum(nil)))
		req.Header.Set("ok-access-timestamp", timestamp)
	}
	resp, err := this.http.Do(req)
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
	if env.Code != "0" {
		return env, fmt.Errorf("OKEx error: %s - %s", env.Code, env.Msg)
	}
	return env, nil
}
