package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Http_test_Post(url string, data interface{}) (content string) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	defer req.Body.Close()
	req.Header.Add("content-type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}
