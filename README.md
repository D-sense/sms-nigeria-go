# sms-nigeria-go
Send SMS to any Nigeria phone number with ease.
---

[![Build Status](https://travis-ci.com/D-sense/sms-nigeria-go.svg?branch=master)](https://travis-ci.com/D-sense/sms-nigeria-go.svg?branch=master)


## Installation
go get github.com/d-sense/sms-nigeria-go


## Usage

> send SMS via Sms_Clone service
```go
// import the package
github.com/d-sense/sms-nigeria-go

// declare and assign notification service
var bulkSmsNigeria sms_nigeria_go.SmsController

//send SMS via Bulk_Sms_Nigeria service
smsCloneData := &sms_nigeria_go.SmsCloneNotification{
	Username:  os.Getenv("USERNAME"),
	Password:  os.Getenv("PASSWORD"),
	Sender:    os.Getenv("SMS_SENDER"),
	Recipient: recipient,
	Message:   textMessage,
}

resultSmsClone, err := bulkSmsNigeria.SendSmsClone(smsCloneData, sms_nigeria_go.SmsCloneNormalRoute)
if err != nil {
	log.Fatal(err)
}
fmt.Println(resultSmsClone)
```

> send SMS via Bulk_Sms_Nigeria service

```go
// import the package
"github.com/d-sense/sms-nigeria-go"


// declare and assign notification service
var smsClone sms_nigeria_go.SmsController

//send SMS via Bulk_Sms_Nigeria service
data := &sms_nigeria_go.BulkSmsNigeriaNotification{
    Sender:    os.Getenv("SMS_SENDER"),
    Recipient: recipient,
		Body:      textMessage,
		ApiToken:  os.Getenv("BULK_SMS_NIGERIA_API_TOKEN"),
	}
}

result, err := smsClone.SendBulkSmsNigeria(data)
if err != nil {
   log.Fatal(err)
}
fmt.Println(result)


// to check credit balance
credential := sms_nigeria_go.SmsCloneCredential{
	Username: os.Getenv("USERNAME"),
	Password: os.Getenv("PASSWORD"),
}
balance, err := repo.CheckBalanceSmsClone(&credential)
if err != nil {
	log.Fatal(err)
}
fmt.Println(balance)
```
---


## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)
- **[MIT license](http://opensource.org/licenses/mit-license.php)**
