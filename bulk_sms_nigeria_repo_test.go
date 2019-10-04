package sms_nigeria_go

import (
	"reflect"
	"strings"
	"testing"
)

var (
	stub StubBulkSmsNigeriaRepository

	apiToken  = "ABCDE"
	sender    = "Adeshina"
	recipient = "09099999"
	body      = "Hello. How are you doing"

	bukSmsNigeriaService = BulkSmsNigeriaComponent{
		&StubBulkSmsNigeriaRepository{},
	}

	bulkSmsNigeriaInfo = BulkSmsNigeriaNotification{
		Sender:    sender,
		Recipient: recipient,
		Body:      "Hello",
		ApiToken:  apiToken,
	}
)

var bulkSmsNigeriaRepo BulkSmsNigeriaRepository
var smsCloneRepo SmsCloneRepository

type StubBulkSmsNigeriaRepository struct{}

func (stub StubBulkSmsNigeriaRepository) BulkSmsNigeria(sms *BulkSmsNigeriaNotification) (BulkSmsNigeriaResponse, error) {
	if result := bulkSmsNigeriaRepo.ValidateBulkSmsNigeriaInput(sms); len(result) > 0 {
		return BulkSmsNigeriaResponse{}, nil
	}

	bulkSmsNigeriaResp := BulkSmsNigeriaResponse{}
	bulkSmsNigeriaResp.Status = 200
	bulkSmsNigeriaResp.Data.Status = "200"
	bulkSmsNigeriaResp.Data.Message = "Success"

	return bulkSmsNigeriaResp, nil
}
func (stub StubBulkSmsNigeriaRepository) ValidateBulkSmsNigeriaInput(sms *BulkSmsNigeriaNotification) (err map[string]interface{}) {
	err = bulkSmsNigeriaRepo.ValidateBulkSmsNigeriaInput(sms)
	return err
}

func TestBulkSmsNigeriaBaseURL(t *testing.T) {

}
func TestContactEndpointFunc(t *testing.T) {

}
func TestBulkSmsNigeriaUrls(t *testing.T) {
	// Testing full URL
	actualFullURL := BulkSmsNigeriaURLCreate + "?api_token=" + apiToken +
		"&from=" + sender + "&to=" + recipient + "&body=" +
		strings.Replace(body, " ", "%20", -1)
	correctFullURL := BulkSmsNigeriaURLCreate + "?api_token=ABCDE&from=Adeshina&to=09099999&body=Hello.%20How%20are%20you%20doing"
	t.Run("returns true on correct full url", func(t *testing.T) {
		result := compareUrls(actualFullURL, correctFullURL)
		assertCorrectURL(t, result, true)
	})

	wrongFullURL := "https://bulksmsnigeria.com/api/v1/sms/create?api_token=/"
	t.Run("returns false on wrong full url", func(t *testing.T) {
		result := compareUrls(actualFullURL, wrongFullURL)
		assertCorrectURL(t, result, false)
	})
}
func TestBulkSmsNigeriaValidation(t *testing.T) {
	//Testing validation
	var validationErr map[string]interface{}
	t.Run("returns true for passed validation", func(t *testing.T) {
		validationErr = stub.ValidateBulkSmsNigeriaInput(&bulkSmsNigeriaInfo)
		assertPassedValidation(t, len(validationErr), 0)
	})
	t.Run("returns false for failed validation", func(t *testing.T) {
		validationErr = stub.ValidateBulkSmsNigeriaInput(&BulkSmsNigeriaNotification{})
		assertFailedValidation(t, len(validationErr), 4)
	})
}
func TestBulkSmsNigeriaHTTPResponse(t *testing.T) {
	bulkSmsNigeriaResp := BulkSmsNigeriaResponse{}
	bulkSmsNigeriaResp.Status = 200
	bulkSmsNigeriaResp.Data.Status = "200"
	bulkSmsNigeriaResp.Data.Message = "Success"

	//Testing endpoint
	t.Run("returns true on non-empty bulkSmsNigeriaResp struct", func(t *testing.T) {
		result, _ := stub.BulkSmsNigeria(&bulkSmsNigeriaInfo)
		assertEndpointCorrectResponseBulkSmsNigeria(t, result, bulkSmsNigeriaResp)
	})

	t.Run("returns true on empty bulkSmsNigeriaResp struct", func(t *testing.T) {
		result, _ := stub.BulkSmsNigeria(&BulkSmsNigeriaNotification{})
		assertEndpointWrongResponseBulkSmsNigeria(t, result, BulkSmsNigeriaResponse{})
	})
}

func compareUrls(urlA, urlB string) bool {
	if urlA == urlB {
		return true
	}
	return false
}

func assertCorrectURL(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("result of url comparison is wrong, got %v want %v", got, want)
	}
}

func assertPassedValidation(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("result of validation is wrong, got %v want %v", got, want)
	}
}

func assertFailedValidation(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("result of validation is wrong, got %v want %v", got, want)
	}
}

func assertEndpointCorrectResponseBulkSmsNigeria(t *testing.T, got, want BulkSmsNigeriaResponse) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("result of full endpoint contact is wrong, got %v want %v", got, want)
	}
}

func assertEndpointWrongResponseBulkSmsNigeria(t *testing.T, got, want BulkSmsNigeriaResponse) {
	t.Helper()
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("result of endpoint result comparison is wrong, got %v want %v", got, want)
	}
}
