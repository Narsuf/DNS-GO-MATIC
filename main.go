package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	userAndPassword := os.Args[1] + ":" + os.Args[2]
	authorization := base64.StdEncoding.EncodeToString([]byte(userAndPassword))
	baseUrl := "https://" + userAndPassword + "@updates.dnsomatic.com/nic/update?hostname=all.dnsomatic.com"
	client := &http.Client{}

	for {
		updateIp(client, authorization, baseUrl)
		time.Sleep(30 * time.Minute)
	}
}

func updateIp(client *http.Client, header string, url string) {
	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		fmt.Println("Request error: ", err)
	}

	req.Header.Set("Authorization", "Basic "+header)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Response error: ", err)
	} else {
		fmt.Println("Response code: ", resp.StatusCode)
		resp.Body.Close()
	}
}
