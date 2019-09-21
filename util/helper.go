package util

import (
	"io/ioutil"
	"net/http"
	"time"
)

func ContactEndpoint(url string) ([]byte, error) {
	spaceClient := http.Client{
		Timeout: time.Second * 2000000, // Maximum of .... secs
	}

	var request *http.Request
	request, err := http.NewRequest(http.MethodGet, url, nil)

	response, err := spaceClient.Do(request)
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return body, err
	}

	//fmt.Println( fmt.Printf("DATA: %s", body))
	return body, err
}
