package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func RandomIntInRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func GetFromAzure(url, token string, v interface{}) {
	log.Println(fmt.Sprintf("Requesting GET from %s", url))

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	sendRequest(req, token, v)

	/*
		// Add authorization header to request
		req.Header.Add("Authorization", "Basic "+basicAuth("", token))
		req.Header.Add("Accept", "application/json;odata=verbose")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Do: ", err)
			return
		}

		defer resp.Body.Close()

		log.Println(fmt.Sprintf("Received response code %d", resp.StatusCode))

		if resp.StatusCode != http.StatusOK {
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		log.Println(fmt.Sprintf("Received response body %s", string(body)))

		if err := json.Unmarshal([]byte(body), &v); err != nil {
			log.Println(err)
		}*/
}

func PostToAzure(url, token, body string, v interface{}) {
	log.Println(fmt.Sprintf("Sending POST to %s", url))

	// Build the request
	bodyReader := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	sendRequest(req, token, v)
}

func sendRequest(req *http.Request, token string, v interface{}) {
	// Add authorization header to request
	req.Header.Add("Authorization", "Basic "+basicAuth("", token))
	req.Header.Add("Accept", "application/json;odata=verbose")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	log.Println(fmt.Sprintf("Received response code %d", resp.StatusCode))

	if resp.StatusCode != http.StatusOK {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(fmt.Sprintf("Received response body %s", string(body)))

	if err := json.Unmarshal([]byte(body), &v); err != nil {
		log.Println(err)
	}
}

// Creates base64 string for PAT network auth
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
