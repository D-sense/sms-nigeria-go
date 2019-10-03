package notification

import (
	"github.com/d-sense/sms-nigeria-go/util"
	"errors"
	"reflect"
	"testing"
	"strings"
)

var (
	stubSmsClone StubSmsCloneRepository

	smsCloneServiceComponent = SmsCloneComponent{
		&StubSmsCloneRepository{},
	}

	smsCloneInfo = SmsCloneNotification{
		Sender:    sender,
		Recipient: recipient,
		Username:  "Hello",
		Password:  "assss",
		Message:   apiToken,
	}

	smsCloneCredentialInfo = SmsCloneCredential{
		Username:    sender,
		Password: recipient,
	}

	creditBalance = "500"
)

type StubSmsCloneRepository struct {}

func (stub StubSmsCloneRepository) SmsClone(sms *SmsCloneNotification, route string) (SmsCloneResponse, error){
	if  result:= stub.ValidateSmsCloneInput(sms); len(result) > 0 {
		return SmsCloneResponse{}, nil
	}

	smsCloneResp := SmsCloneResponse{}
	smsCloneResp.BatchCode = "1234"
	smsCloneResp.BatchDescription = "Processed"
	smsCloneResp.StatusCode = "200"
	smsCloneResp.Recipient = "Adeshina"
	smsCloneResp.MessageID = "01234"
	smsCloneResp.MessageStatus = "Success"
	smsCloneResp.StatusDescription = "Sent"

	return smsCloneResp, nil
}
func (stub StubSmsCloneRepository) SmsCloneCheckBalance(sms *SmsCloneCredential) (response string, err error){
	if  result:= stub.ValidateSmsCloneCredentials(sms); len(result) > 0 {
		return response, errors.New("validation errors")
	}

	response = creditBalance
	return response, nil
}
func (stub StubSmsCloneRepository)	ValidateSmsCloneInput(smsInfo *SmsCloneNotification) (err map[string]interface{}){
	err = smsCloneRepo.ValidateSmsCloneInput(smsInfo)
	return err
}
func (stub StubSmsCloneRepository)	ValidateSmsCloneCredentials(smsInfo *SmsCloneCredential) (err map[string]interface{}){
	err = smsCloneRepo.ValidateSmsCloneCredentials(smsInfo)
	return err
}

func TestSmsCloneUrls (t *testing.T){
	// Testing full URL
	actualFullURL := util.SmsCloneNormalRouteURLCreate+"?api_token="+ apiToken +
		"&from="+ sender +"&to="+ recipient +"&body=" +
		strings.Replace(body, " ", "%20", -1)
	correctFullURL := util.SmsCloneNormalRouteURLCreate+"?api_token=ABCDE&from=Adeshina&to=09099999&body=Hello.%20How%20are%20you%20doing"
	t.Run("returns TRUE on correct full url", func(t *testing.T){
		result := compareUrls(actualFullURL, correctFullURL)
		assertCorrectURL(t, result, true)
	})

	wrongFullURL := "https://smsclone.com/api/v1/sms/create?api_token=/"
	t.Run("returns FALSE on wrong full url", func(t *testing.T){
		result := compareUrls(actualFullURL, wrongFullURL)
		assertCorrectURL(t, result, false)
	})
}
func TestSmsCloneValidation (t *testing.T){
	//Testing validation
	var validationErr map[string]interface{}
	t.Run("returns true for passed validation", func(t *testing.T){
		validationErr = smsCloneServiceComponent.NotifySmsClone.ValidateSmsCloneInput(&smsCloneInfo)
		assertPassedValidation(t, len(validationErr), 0)
	})

	t.Run("returns true for failed validation", func(t *testing.T){
		validationErr = smsCloneServiceComponent.NotifySmsClone.ValidateSmsCloneInput(&SmsCloneNotification{})
		assertFailedValidation(t, len(validationErr), 5)
	})
}
func TestSmsCloneCheckBalance (t *testing.T){
	//Testing account balance
	t.Run("returns true for passed check balance", func(t *testing.T){
		result, _ := smsCloneServiceComponent.NotifySmsClone.SmsCloneCheckBalance(&smsCloneCredentialInfo)
		assertCheckBalanceSmsClone(t, result, creditBalance)
	})

	t.Run("returns true for failed balance check", func(t *testing.T){
		result, _ := smsCloneServiceComponent.NotifySmsClone.SmsCloneCheckBalance(&SmsCloneCredential{})
		assertCheckBalanceSmsClone(t, result, "")
	})
}
func TestSmsCloneHTTPResponse (t *testing.T) {
	smsCloneResp := SmsCloneResponse{}
	smsCloneResp.BatchCode = "1234"
	smsCloneResp.BatchDescription = "Processed"
	smsCloneResp.StatusCode = "200"
	smsCloneResp.Recipient = "Adeshina"
	smsCloneResp.MessageID = "01234"
	smsCloneResp.MessageStatus = "Success"
	smsCloneResp.StatusDescription = "Sent"

	//Testing endpoint
	t.Run("returns true on non-empty bulkSmsNigeriaResp struct", func(t *testing.T){
		result, _ := smsCloneServiceComponent.NotifySmsClone.SmsClone(&smsCloneInfo, "")
		assertEndpointCorrectResponseSmsClone(t, result, smsCloneResp)
	})

	t.Run("returns true on empty bulkSmsNigeriaResp struct", func(t *testing.T){
		result, _ := smsCloneServiceComponent.NotifySmsClone.SmsClone(&SmsCloneNotification{}, "")
		assertEndpointWrongResponseSmsClone(t, result, SmsCloneResponse{})
	})
}

func assertEndpointCorrectResponseSmsClone(t *testing.T, got, want SmsCloneResponse) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("result of full endpoint contact is wrong, got %v want %v", got, want)
	}
}
func assertCheckBalanceSmsClone(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("result of balance check is wrong, got %v want %v", got, want)
	}
}
func assertEndpointWrongResponseSmsClone(t *testing.T, got, want SmsCloneResponse) {
	t.Helper()
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("result of endpoint result comparison is wrong, got %v want %v", got, want)
	}
}





