package notification

import (
	"fmt"
	"errors"
)

type SmsController struct {
	bulkSmsNigeriaNotificationService
	smsCloneNotificationService
}

func (SmsController) SendBulkSmsNigeria(sms *BulkSmsNigeriaNotification) (response bulkSmsNigeriaResponse, err error ) {
	smsRepository := bulkSmsNigeriaRepository{}
	smsValidator := bulkSmsNigeriaValidator{}

	smsServiceComponent := bulkSmsNigeriaNotificationService{
		&smsRepository,
			&smsValidator,
	}

	var validationErr map[string]interface{}
	validationErr = smsServiceComponent.validate.ValidateBulkSmsNigeriaInput(sms)
	if len(validationErr) > 0  {
		err = errors.New(fmt.Sprintf("%v", validationErr))
		return response, err
	}

	response, err = smsServiceComponent.smsNotificationRepo.BulkSmsNigeria(sms)
	if err != nil  {
		return response, err
	}

	return response, nil
}

func (SmsController) SendSmsClone(sms *SmsCloneNotification, route string) (response smsCloneResponse, err error ) {
	smsRepository := smsCloneRepository{}
	smsValidator := SmsCloneValidator{}

	smsServiceComponent := smsCloneNotificationService{
		&smsRepository,
			&smsValidator,
	}

	var validationErr map[string]interface{}
	validationErr = smsServiceComponent.validate.ValidateSmsCloneInput(sms)
	if len(validationErr) > 0  {
		err = errors.New(fmt.Sprintf("%v", validationErr))
		return response, err
	}

	response, err = smsServiceComponent.smsNotificationRepo.SmsClone(sms, route)
	if err != nil  {
		return response, err
	}

	return response, nil
}

func (SmsController) CheckBalanceSmsClone(sms *SmsCloneNotification) (response string, err error ) {
	smsRepository := smsCloneRepository{}
	smsValidator := SmsCloneValidator{}

	smsServiceComponent := smsCloneNotificationService{
		&smsRepository,
			&smsValidator,
	}

	var validationErr map[string]interface{}
	validationErr = smsServiceComponent.validate.ValidateSmsCloneCredentials(sms)
	if len(validationErr) > 0  {
		err = errors.New(fmt.Sprintf("%v", validationErr))
		return response, err
	}

	response, err = smsServiceComponent.smsNotificationRepo.SmsCloneCheckBalance()
	if err != nil  {
		return response, err
	}

	return response, nil
}

