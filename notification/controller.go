package notification

import (
	"fmt"
	"errors"
)

type SmsController struct {
	smsNotificationService
}

func (SmsController) SendBulkSmsNigeria(sms *BulkSmsNigeriaNotification) (response bulkSmsNigeriaResponse, err error ) {
	smsRepository := bulkSmsNigeriaRepository{}
	smsValidator := smsValidator{}

	smsServiceComponent := smsNotificationService{
		&smsRepository,
			&smsValidator,
	}

	var validationErr map[string]interface{}
	validationErr = smsServiceComponent.validate.ValidateBulkSmsNigeriaInput(sms)
	if len(validationErr) > 0  {
		err = errors.New(fmt.Sprintf("%v", validationErr))
		return response, err
	}

	response, err = smsServiceComponent.smsNotificationRepo.SendSMS(sms)
	if err != nil  {
		return response, err
	}

	return response, nil
}
