# sms-nigeria-go — Send SMS to any Nigeria phone-number with ease
--

[![Build Status](https://travis-ci.com/D-sense/sms-nigeria-go.svg?branch=master)](https://travis-ci.com/D-sense/sms-nigeria-go.svg?branch=master)


## Installation
go get github.com/d-sense/sms-nigeria-go


## Usage

> send SMS via Bulk_Sms_Nigeria service (first, open and redeem an account here to obtain credentials: https://www.bulksmsnigeria.com/)

```go
// import the package
github.com/d-sense/sms-nigeria-go

// declare the notification service
var bulkSmsNigeria sms_nigeria_go.SmsController

// declare, initialize, and send data
data := &sms_nigeria_go.BulkSmsNigeriaNotification{
	Sender:    os.Getenv("SMS_SENDER"),
	Recipient: recipient,
	Body:      textMessage,
	ApiToken:  os.Getenv("BULK_SMS_NIGERIA_API_TOKEN"),
}

result, err := repo.SendBulkSmsNigeria(data)
if err != nil {
	log.Fatal(err)
}
fmt.Println(result)
```


> send SMS via Sms_Cone service (first, open and redeem an account here to obtain credentials: http://smsclone.com)

```go
// import the package
github.com/d-sense/sms-nigeria-go

// declare the notification service
var smsClone sms_nigeria_go.SmsController

// declare, initialize, and send data
smsCloneData := &sms_nigeria_go.SmsCloneNotification{
	Username:  os.Getenv("USERNAME"),
	Password:  os.Getenv("PASSWORD"),
	Sender:    os.Getenv("SMS_SENDER"),
	Recipient: recipient,
	Message:   textMessage,
}

result, err := repo.SendSmsClone(smsCloneData, sms_nigeria_go.SmsCloneNormalRoute)
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

#### Note that for Sms_Clone service, you can use different methods to send a SMS and ensure the message will be delivered. In my experience, "dnd-fallback" seems like the most reliable (charge varies per method though).
Method | Constant
:----: | --------
sendsms route | SmsCloneNormalRoute
dnd-route | SmsCloneDndRoute
dnd-fallback | SmsCloneNormalAndDndRoute
---


## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)
- **[MIT license](http://opensource.org/licenses/mit-license.php)**
