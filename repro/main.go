package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var (
	reproURL = "https://marketing.reproio.com/v1/push/x7y15pvx/deliver"
	reproToken = "f8694450-fb2e-4209-834f-b02b763d6251"
)

func main() {
	userIDs := []string{
		"ap-northeast-1:79f83dff-9c07-4a56-885c-90b52230e7d3",
	}
	reproURL, err := url.Parse(reproURL)
	if err != nil {
		spew.Dump(err)
		return
	}

	body := map[string]interface{}{
		"audience": map[string][]string{
			"user_ids": userIDs,
		},
	}

	values, err := json.Marshal(body)
	if err != nil {
		spew.Dump(err)
		return
	}

	req, err := http.NewRequest("POST", reproURL.String(), bytes.NewBuffer(values))
	if err != nil {
		spew.Dump(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Repro-Token", reproToken)
	resp, err := (&http.Client{Timeout: 5 * time.Second}).Do(req)
	if err != nil {
		spew.Dump(err)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		spew.Dump(err)
		return
	}

	var result interface{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		spew.Dump(err)
		return
	}

	spew.Dump()
	return
}
