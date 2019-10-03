package notification

import (
	"encoding/json"
	"github.com/d-sense/go-sms-nigeria/util"
	"os"
	"strings"
)

type BulkSmsNigeriaNotification struct {
	Sender       string
	Recipient    string
	Body         string
	ApiToken     string
}

type BulkSmsNigeriaResponse struct {
	Status int `json:"0"`
	Data struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"data"`
}

type BulkSmsNigeriaRepository struct{}

func (*BulkSmsNigeriaRepository) BulkSmsNigeria(sms *BulkSmsNigeriaNotification) (response BulkSmsNigeriaResponse, err error) {
	var smsData BulkSmsNigeriaNotification
	var bulkSms BulkSmsNigeriaResponse

	smsData.Sender = sms.Sender
	smsData.Recipient = sms.Recipient
	smsData.Body = sms.Body

	prepareURL := util.BulkSmsNigeriaURLCreate+"?api_token="+os.Getenv("BULK_SMS_NIGERIA_API_TOKEN")+
		"&from="+sms.Sender+"&to="+sms.Recipient+"&body=" +
		strings.Replace(sms.Body, " ", "%20", -1)

	// contact endpoint
	var result []byte
	result, err = util.ContactEndpoint(prepareURL)
	if err != nil  {
		return BulkSmsNigeriaResponse{}, err
	}

	err = json.Unmarshal(result, &bulkSms)

	return bulkSms, err
}

func (*BulkSmsNigeriaRepository) ValidateBulkSmsNigeriaInput(smsInfo *BulkSmsNigeriaNotification) (err map[string]interface{}) {
	err = make(map[string]interface{})

	if smsInfo.Sender == "" {
		err["Sender"] = util.ErrMissingSender
	}

	if smsInfo.Recipient == "" {
		err["Recipient"] = util.ErrMissingRecipient
	}

	if smsInfo.Body == "" {
		err["Body"] = util.ErrMissingBody
	}

	if smsInfo.ApiToken == "" {
		err["ApiToken"] = util.ErrMissingApiToken
	}

	return
}
