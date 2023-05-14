package main

import (
	"bytes"
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
		user := []byte("{\"Apartment\":\"string\",\"Balcony\":true,\"City\":\"string\",\"Country\":\"string\",\"Garage\":true,\"Garden\":true,\"HeatType\":0,\"Housenumber\":\"string\",\"Id\":\"string\",\"Landlord\":\"string\",\"Rooms\":0,\"Size\":0,\"Status\":0,\"Street\":\"string\",\"Tenant\":\"string\",\"Type\":0}")
		Request("POST", user, "/user")
	}
}

func BenchmarkPutRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := []byte("{\"Apartment\":\"string\",\"Balcony\":true,\"City\":\"string\",\"Country\":\"string\",\"Garage\":true,\"Garden\":true,\"HeatType\":0,\"Housenumber\":\"string\",\"Id\":\"string\",\"Landlord\":\"string\",\"Rooms\":0,\"Size\":0,\"Status\":0,\"Street\":\"string\",\"Tenant\":\"string\",\"Type\":0}")
		Request("PUT", user, "/user/"+uuid.New().String())
	}
}

func Request(typ string, body []byte, route string) error {

	req, err := http.NewRequest(typ, "http://127.0.0.1:3000"+route, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
