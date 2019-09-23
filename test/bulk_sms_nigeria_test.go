package test

import (
	"github.com/d-sense/go-sms-nigeria/util"
	"strings"
	"testing"
)

func TestBulkSmsNigeriaUrls (t *testing.T){
	// Testing base URL
	correctBaseURL := "https://bulksmsnigeria.com/api/v1/sms/create"
	t.Run("returns TRUE on correct base url", func(t *testing.T){
		result := compareUrls(util.BulkSmsNigeriaURLCreate, correctBaseURL)
		assertResponseBody(t, result, true)
	})

	wrongBaseURL := "this/is/a/wrong/url"
	t.Run("returns FALSE on wrong base url", func(t *testing.T){
		result := compareUrls(util.BulkSmsNigeriaURLCreate, wrongBaseURL)
		assertResponseBody(t, result, false)
	})

	// Testing full URL
	ApiToken :="ABCDE"
	sender := "Adeshina"
	recipient := "09099999"
	body := "Hello. How are you doing"
	actualFullURL := util.BulkSmsNigeriaURLCreate+"?api_token="+ApiToken+
		"&from="+sender+"&to="+recipient+"&body=" +
		strings.Replace(body, " ", "%20", -1)
	correctFullURL := util.BulkSmsNigeriaURLCreate+"?api_token=ABCDE&from=Adeshina&to=09099999&body=Hello.%20How%20are%20you%20doing"

	t.Run("returns TRUE on correct full url", func(t *testing.T){
		result := compareUrls(actualFullURL, correctFullURL)
		assertResponseBody(t, result, true)
	})

	wrongFullURL := "https://bulksmsnigeria.com/api/v1/sms/create?api_token=/"
	t.Run("returns FALSE on wrong full url", func(t *testing.T){
		result := compareUrls(actualFullURL, wrongFullURL)
		assertResponseBody(t, result, false)
	})
}


type bulkSmsNigeriaResponse struct {
	Status int `json:"0"`
	Data struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"data"`
}

type bulkSmsNigeriaNotification struct {
	Sender       string
	Recipient    string
	Body         string
	ApiToken     string
}

//type stubBulkSmsNigeriaRepository struct {}
//func (stub stubBulkSmsNigeriaRepository) BulkSmsNigeria(sms *bulkSmsNigeriaNotification) (bulkSmsNigeriaResponse, error){
//    return bulkSmsNigeriaResponse{}, nil
//}
//
//type stubBulkSmsNigeriaRepository struct {}
//func (stub stubBulkSmsNigeriaRepository) BulkSmsNigeria(sms *bulkSmsNigeriaNotification) (bulkSmsNigeriaResponse, error){
//	return bulkSmsNigeriaResponse{}, nil
//}
//
//
//func TestHTTClient (t *testing.T) {
//    t.Run("", func(t *testing.T){
//
//	})
//}





func compareUrls(urlA, urlB string) bool {
	if urlA == urlB {
		return true
	}
	return false
}

func assertResponseBody(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("result of url comparison is wrong, got %v want %v", got, want)
	}
}