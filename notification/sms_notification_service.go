package notification

type smsNotificationRepository interface {
	SendSMS(sms *BulkSmsNigeriaNotification) (bulkSmsNigeriaResponse, error)
}

type validateSmsNotification interface {
	ValidateBulkSmsNigeriaInput(smsInfo *BulkSmsNigeriaNotification) (err map[string]interface{})
}

type smsNotificationService struct {
	smsNotificationRepo smsNotificationRepository
	validate validateSmsNotification
}