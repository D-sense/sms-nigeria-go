package notification

//Bulk_SMS_Nigeria service
type smsNotificationRepository interface {
	BulkSmsNigeria(sms *BulkSmsNigeriaNotification) (bulkSmsNigeriaResponse, error)
	ValidateBulkSmsNigeriaInput(smsInfo *BulkSmsNigeriaNotification) (err map[string]interface{})
}

type smsCloneNotificationRepository interface {
	SmsClone(sms *SmsCloneNotification, route string) (smsCloneResponse, error)
	SmsCloneCheckBalance() (response string, err error)
	ValidateSmsCloneInput(smsInfo *SmsCloneNotification) (err map[string]interface{})
	ValidateSmsCloneCredentials(smsInfo *SmsCloneNotification) (err map[string]interface{})
}

type bulkSmsNigeriaNotificationService struct {
	smsNotificationRepo smsNotificationRepository
}

type smsCloneNotificationService struct {
	smsNotificationRepo smsCloneNotificationRepository
}