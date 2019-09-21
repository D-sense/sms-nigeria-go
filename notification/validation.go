package notification

import (
	"github.com/d-sense/go-sms-nigeria/util"
)

type smsValidator struct{}

func (*smsValidator) ValidateBulkSmsNigeriaInput(smsInfo *BulkSmsNigeriaNotification) (err map[string]interface{}) {
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
