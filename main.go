package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	updateURL      = "https://updates.dnsomatic.com/nic/update?hostname=all.dnsomatic.com"
	updateInterval = 30 * time.Minute
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: dns-go-matic <username> <password>")
		os.Exit(1)
	}

	username := os.Args[1]
	password := os.Args[2]

	credentials := username + ":" + password
	authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials))

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	log.Printf("Starting DNS-O-Matic updater (interval: %v)", updateInterval)

	for {
		if err := updateIP(client, authorization); err != nil {
			log.Printf("Error updating IP: %v", err)
		}
		time.Sleep(updateInterval)
	}
}

func updateIP(client *http.Client, authorization string) error {
	req, err := http.NewRequest("GET", updateURL, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", authorization)
	req.Header.Set("User-Agent", "DNS-GO-MATIC/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	log.Printf("Response [%d]: %s", resp.StatusCode, string(body))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
