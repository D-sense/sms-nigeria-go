package sms_nigeria_go

import (
	"encoding/json"
	"strings"
)

type BulkSmsNigeriaNotification struct {
	Sender    string
	Recipient string
	Body      string
	ApiToken  string
}

type BulkSmsNigeriaResponse struct {
	Status int `json:"0"`
	Data   struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"data"`
}

type BulkSmsNigeriaRepository struct{}

func (*BulkSmsNigeriaRepository) BulkSmsNigeria(sms *BulkSmsNigeriaNotification) (response BulkSmsNigeriaResponse, err error) {
	prepareURL := BulkSmsNigeriaURLCreate + "?api_token=" +sms.ApiToken+
		"&from=" + sms.Sender + "&to=" + sms.Recipient + "&body=" +
		strings.Replace(sms.Body, " ", "%20", -1)

	// contact endpoint
	var result []byte
	result, err = ContactEndpoint(prepareURL)
	if err != nil {
		return BulkSmsNigeriaResponse{}, err
	}

	err = json.Unmarshal(result, &response)

	return response, err
}

func (*BulkSmsNigeriaRepository) ValidateBulkSmsNigeriaInput(smsInfo *BulkSmsNigeriaNotification) (err map[string]interface{}) {
	err = make(map[string]interface{})

	if smsInfo.Sender == "" {
		err["Sender"] = ErrMissingSender
	}

	if smsInfo.Recipient == "" {
		err["Recipient"] = ErrMissingRecipient
	}

	if smsInfo.Body == "" {
		err["Body"] = ErrMissingBody
	}

	if smsInfo.ApiToken == "" {
		err["ApiToken"] = ErrMissingApiToken
	}

	return
}
