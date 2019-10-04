package sms_nigeria_go

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func ContactEndpoint(url string) ([]byte, error) {
	if result := isConnected(); result == false {
		response := make([]byte, 1)
		return response, errors.New(InternetConnectionErr)
	}

	spaceClient := http.Client{
		Timeout: time.Second * 2000000, // Maximum of .... secs
	}

	var request *http.Request
	request, err := http.NewRequest(http.MethodGet, url, nil)

	response, err := spaceClient.Do(request)
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		err = readErr
		return body, err
	}

	return body, nil
}

func isConnected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
