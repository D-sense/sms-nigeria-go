package notification

import (
	"github.com/d-sense/go-sms-nigeria/util"
)

type bulkSmsNigeriaValidator struct{}
type SmsCloneValidator struct{}

func (*bulkSmsNigeriaValidator) ValidateBulkSmsNigeriaInput(smsInfo *BulkSmsNigeriaNotification) (err map[string]interface{}) {
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

func (*SmsCloneValidator) ValidateSmsCloneInput(smsInfo *SmsCloneNotification) (err map[string]interface{}) {
	err = make(map[string]interface{})

	if smsInfo.Username == "" {
		err["Username"] = util.ErrMissingUsername
	}

	if smsInfo.Sender == "" {
		err["Sender"] = util.ErrMissingSender
	}

	if smsInfo.Recipient == "" {
		err["Recipient"] = util.ErrMissingRecipient
	}

	if smsInfo.Message == "" {
		err["Message"] = util.ErrMissingMessage
	}

	if smsInfo.Password == "" {
		err["Password"] = util.ErrMissingPassword
	}

	return
}

func (*SmsCloneValidator) ValidateSmsCloneCredentials(smsInfo *SmsCloneNotification) (err map[string]interface{}) {
	err = make(map[string]interface{})

	if smsInfo.Username == "" {
		err["Username"] = util.ErrMissingUsername
	}

	if smsInfo.Password == "" {
		err["Password"] = util.ErrMissingPassword
	}

	return
}
