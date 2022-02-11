package main

import (
	"bytes"
	"log"
	"net/http"
)

func MmSend(url, message string) {
	var body = []byte(`{"text":"` + message + `"}`)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
}
