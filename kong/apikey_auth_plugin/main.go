package main

import (
	"io"
	"net/http"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

type Config struct {
	ApiKeyVerificationURL string `json:"apikey_verification_url"`
}

func New() interface{} {
	return &Config{}
}

func (conf *Config) Access(kong *pdk.PDK) {
	apiKey, err := kong.Request.GetHeader("x-api-key")
	if err != nil {
		kong.Response.Exit(401, []byte(err.Error()), nil)
		return
	}

	req, err := http.NewRequest("GET", conf.ApiKeyVerificationURL, nil)
	if err != nil {
		kong.Log.Err(err.Error())
		kong.Response.Exit(500, []byte(err.Error()), nil)
		return
	}

	req.Header.Set("x-api-key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		kong.Log.Err(err.Error())
		kong.Response.Exit(500, []byte(err.Error()), nil)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		kong.Log.Err(string(body))
		kong.Response.Exit(resp.StatusCode, body, nil)
		return
	}
}

func main() {
	server.StartServer(New, "0.1", 1000)
}
