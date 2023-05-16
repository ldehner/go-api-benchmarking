package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// This is the main function
	NewRequest("POST", []byte(`{"id":"string","firstName":"string","lastName":"string","phone":"string","mail":"string","country":"string","city":"string","street":"string","housenumber":"string","apartment":"string"}`), "/user")
}

func NewRequest(typ string, body []byte, route string) error {

	req, err := http.NewRequest(typ, "http://127.0.0.1:8080"+route, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Accept-Encoding", "identity") // NOTE THIS LINE
	req.Header.Add("Content-Type", "application/json")
	req.Close = true
	client := &http.Client{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // <--- Problem
	}
	client.Transport = tr
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("here")
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	return nil
}

func NewRequest2(body []byte, route string) error {
	resp, err1 := http.Get("http://127.0.0.1:8080" + route)

	if err1 != nil {
		log.Printf("Request Failed: %s", err1)
		return err1
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return err
	} // Log the request body
	bodyString := string(body)
	log.Print(bodyString)
	return nil
}
