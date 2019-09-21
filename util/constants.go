package util

const (
	//Bulk_SMS_Nigeria
	BulkSmsNigeriaURLCreate = "https://bulksmsnigeria.com/api/v1/sms/create"

	//validation
	ErrMissingSender = "Sender field is missing/empty"
	ErrMissingRecipient = "Recipient field is missing/empty"
	ErrMissingBody = "Body field is missing/empty"
	ErrMissingApiToken = "API_TOKEN field is missing/empty"

    ThirdPartyErr = "external_service_response_error"
)