package requests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Post(url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("could not send the request to the server: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get the response from the server: %s", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body: %s", err)
	}

	return body, nil
}

func Get_Authorized(url string, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %s", err)
	}

	req.Header.Add("Accept", "application/vnd.api+json, application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get the response from the server: %s", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body: %s", err)
	}

	return body, nil
}
