package dnsomatic

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func UpdateIP(user, password string) error {
	userPassword := fmt.Sprintf("%s:%s", user, password)

	url := fmt.Sprintf("https://%s@updates.dnsomatic.com/nic/update?hostname=all.dnsomatic.com", userPassword)

	// both GET and POST work. They recommend we use GET
	// @see https://dnsomatic.com/docs/api
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("error creating request", err.Error())
		return err
	}

	authorization := base64.StdEncoding.EncodeToString([]byte(userPassword))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s, ", authorization))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error making request", err.Error())
		return err
	}

	// check response
	// @see https://dnsomatic.com/docs/api
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if strings.HasPrefix(string(body), "good ") || strings.HasPrefix(string(body), "nochg ") {
		return nil
	}

	err = fmt.Errorf("error updating ip %s", string(body))
	log.Println(err)
	return err
}
