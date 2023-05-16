package main

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func BenchmarkGetRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Request("GET", nil, "/user/"+uuid.New().String())
	}
}

func BenchmarkPostRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := []byte(`{"id":"string","firstName":"string","lastName":"string","phone":"string","mail":"string","country":"string","city":"string","street":"string","housenumber":"string","apartment":"string"}`)
		Request("POST", user, "/user")
	}
}

func BenchmarkPutRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := []byte(`{"id":"string","firstName":"string","lastName":"string","phone":"string","mail":"string","country":"string","city":"string","street":"string","housenumber":"string","apartment":"string"}`)
		Request("PUT", user, "/user/"+uuid.New().String())
	}
}

func Request(typ string, body []byte, route string) error {
	req, err := http.NewRequest(typ, "http://127.0.0.1:8080"+route, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Accept-Encoding", "identity") // NOTE THIS LINE
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // <--- Problem
	}
	client.Transport = tr
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
