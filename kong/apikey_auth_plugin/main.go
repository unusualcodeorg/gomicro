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
	headers := map[string][]string{"Content-Type": {"application/json"}}

	apiKey, err := kong.Request.GetHeader("x-api-key")
	if err != nil {
		kong.Response.Exit(401, []byte(err.Error()), headers)
		return
	}

	req, err := http.NewRequest("GET", conf.ApiKeyVerificationURL, nil)
	if err != nil {
		kong.Log.Err(err.Error())
		kong.Response.Exit(500, []byte(err.Error()), headers)
		return
	}

	req.Header.Set("x-api-key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		kong.Log.Err(err.Error())
		kong.Response.Exit(500, []byte(err.Error()), headers)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		kong.Log.Err(string(body))
		kong.Response.Exit(resp.StatusCode, body, headers)
		return
	}
}

func main() {
	server.StartServer(New, "0.1", 1000)
}
