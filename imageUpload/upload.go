package main

import (
	"fmt"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var reqBody io.Reader
	reqBody, _ = os.Open("/home/kaikun/Pictures/yingyang.png")
	reqURL := "http://localhost:4950/"

	/*
		resp, _ := http.Post(, "image/png", reader)
		fmt.Printf("Response length : %d\n", resp.ContentLength)
	*/
	client := &http.Client{}

	req, err := http.NewRequest("POST", reqURL, reqBody)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}
	req.Header.Add("Content-Length", "1076")
	req.Header.Del("Transfer-Encoding")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}

	fmt.Printf("Response: %s", body)
}
