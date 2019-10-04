package sms_nigeria_go

const (
	//Bulk_SMS_Nigeria
	BulkSmsNigeriaURLCreate = "https://bulksmsnigeria.com/api/v1/sms/create"

	//SMS_Clone
	SmsCloneNormalRouteURLCreate       = "https://smsclone.com/api/sms/sendsms"
	SmsCloneDndRouteURLCreate          = "https://smsclone.com/api/sms/dnd-route"
	SmsCloneNormalAndDndRouteURLCreate = "https://smsclone.com/api/sms/dnd-fallback"
	SmsCloneCheckBalanceURL            = "https://smsclone.com/api/sms/balance"

	SmsCloneNormalRoute       = "normal_route"
	SmsCloneDndRoute          = "dnd_route"
	SmsCloneNormalAndDndRoute = "normal_dnd_route"
	PossibleCredentialsErr    = "could not authenticate your credentials (username and/or password)"
	InternetConnectionErr     = "you are not connected to the internet"

	//validation
	ErrMissingSender    = " Sender field is missing/empty."
	ErrMissingUsername  = " Username field is missing/empty."
	ErrMissingRecipient = " Recipient field is missing/empty."
	ErrMissingBody      = " Body field is missing/empty."
	ErrMissingMessage   = " Message field is missing/empty."
	ErrMissingApiToken  = " API_TOKEN field is missing/empty."
	ErrMissingPassword  = " Password field is missing/empty."
)
