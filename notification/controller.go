package notification

import (
	"fmt"
	"errors"
)

type SmsController struct {
	BulkSmsNigeriaComponent
	SmsCloneComponent
}

func (SmsController) SendBulkSmsNigeria(sms *BulkSmsNigeriaNotification) (response BulkSmsNigeriaResponse, err error ) {
	var smsRepository BulkSmsNigeriaRepository

	smsServiceComponent := BulkSmsNigeriaComponent{
		&smsRepository,
	}

	var validationErr map[string]interface{}
	validationErr = smsServiceComponent.NotifyBulkSmsNigeria.ValidateBulkSmsNigeriaInput(sms)
	if len(validationErr) > 0  {
		err = errors.New(fmt.Sprintf("%v", validationErr))
		return response, err
	}

	response, err = smsServiceComponent.NotifyBulkSmsNigeria.BulkSmsNigeria(sms)
	if err != nil  {
		return response, err
	}

	return response, nil
}

func (SmsController) SendSmsClone(sms *SmsCloneNotification, route string) (response SmsCloneResponse, err error ) {
	smsRepository := SmsCloneRepository{}

	smsServiceComponent := SmsCloneComponent{
		&smsRepository,
	}

	var validationErr map[string]interface{}
	validationErr = smsServiceComponent.NotifySmsClone.ValidateSmsCloneInput(sms)
	if len(validationErr) > 0  {
		err = errors.New(fmt.Sprintf("%v", validationErr))
		return response, err
	}

	response, err = smsServiceComponent.NotifySmsClone.SmsClone(sms, route)
	if err != nil  {
		return response, err
	}

	return response, nil
}

func (SmsController) CheckBalanceSmsClone(sms *SmsCloneCredential) (response string, err error ) {
	smsRepository := SmsCloneRepository{}

	smsServiceComponent := SmsCloneComponent{
		&smsRepository,
	}

	var validationErr map[string]interface{}
	validationErr = smsServiceComponent.NotifySmsClone.ValidateSmsCloneCredentials(sms)
	if len(validationErr) > 0  {
		err = errors.New(fmt.Sprintf("%v", validationErr))
		return response, err
	}

	response, err = smsServiceComponent.NotifySmsClone.SmsCloneCheckBalance(sms)
	if err != nil  {
		return response, err
	}

	return response, nil
}

