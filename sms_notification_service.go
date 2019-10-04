package sms_nigeria_go

//Bulk_SMS_Nigeria service
type BulkSmsNigeriaRepositoryInterface interface {
	BulkSmsNigeria(sms *BulkSmsNigeriaNotification) (BulkSmsNigeriaResponse, error)
	ValidateBulkSmsNigeriaInput(smsInfo *BulkSmsNigeriaNotification) (err map[string]interface{})
}

type SmsCloneNotificationRepositoryInterface interface {
	SmsClone(sms *SmsCloneNotification, route string) (SmsCloneResponse, error)
	SmsCloneCheckBalance(sms *SmsCloneCredential) (response string, err error)
	ValidateSmsCloneInput(smsInfo *SmsCloneNotification) (err map[string]interface{})
	ValidateSmsCloneCredentials(smsInfo *SmsCloneCredential) (err map[string]interface{})
}

type BulkSmsNigeriaComponent struct {
	NotifyBulkSmsNigeria BulkSmsNigeriaRepositoryInterface
}

type SmsCloneComponent struct {
	NotifySmsClone SmsCloneNotificationRepositoryInterface
}
