package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	cert, err := tls.LoadX509KeyPair("./cert/server.pem", "./cert/server.key")

	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("read err: %s", err)
	}

	fmt.Printf("response %d: %s\n", resp.StatusCode, string(body))
}
