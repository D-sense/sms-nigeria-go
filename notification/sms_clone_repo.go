package notification

import (
	"errors"
	"os"
	"strings"
	"github.com/d-sense/sms-nigeria-go/util"
)

type SmsCloneNotification struct {
	Username     string
	Password     string
	Sender       string
	Recipient    string
	Message      string
}

type SmsCloneCredential struct {
	Username     string
	Password     string
}

type SmsCloneResponse struct {
	BatchCode                      string                 
	BatchDescription               string                 
	StatusCode                     string                 
	Recipient                      string                 
	MessageID                      string                 
	MessageStatus                  string                 
	StatusDescription              string
}

type SmsCloneRepository struct{}

func (*SmsCloneRepository) SmsClone(sms *SmsCloneNotification, route string) (response SmsCloneResponse, err error) {
	var smsData SmsCloneNotification
	var smsClone SmsCloneResponse
    var preparedURL string

	smsData.Sender = sms.Sender
	smsData.Recipient = sms.Recipient
	smsData.Message = sms.Message

	switch route {
	case util.SmsCloneNormalRoute :
		preparedURL = util.SmsCloneNormalRouteURLCreate+"?username="+os.Getenv("USERNAME")+
			"&password="+os.Getenv("PASSWORD")+"&sender="+sms.Sender+"&recipient="+sms.Recipient+"&message=" +
			strings.Replace(sms.Message, " ", "%20", -1)

	case util.SmsCloneDndRoute :
		preparedURL = util.SmsCloneDndRouteURLCreate+"?username="+os.Getenv("USERNAME")+
			"&password="+os.Getenv("PASSWORD")+"&sender="+sms.Sender+"&recipient="+sms.Recipient+"&message=" +
			strings.Replace(sms.Message, " ", "%20", -1)

	case util.SmsCloneNormalAndDndRoute :
		preparedURL = util.SmsCloneNormalAndDndRouteURLCreate+"?username="+os.Getenv("USERNAME")+
			"&password="+os.Getenv("PASSWORD")+"&sender="+sms.Sender+"&recipient="+sms.Recipient+"&message=" +
			strings.Replace(sms.Message, " ", "%20", -1)

	default :
		preparedURL = util.SmsCloneNormalRouteURLCreate+"?username="+os.Getenv("USERNAME")+
			"&password="+os.Getenv("PASSWORD")+"&sender="+sms.Sender+"&recipient="+sms.Recipient+"&message=" +
			strings.Replace(sms.Message, " ", "%20", -1)
	}

	// contact endpoint
	var result []byte
	result, err = util.ContactEndpoint(preparedURL)
	if err != nil  {
		return SmsCloneResponse{}, err
	}

	smsClone, err = ResponseParser(string(result))
	if err != nil  {
		return SmsCloneResponse{}, err
	}

	return smsClone, err
}

func (*SmsCloneRepository) SmsCloneCheckBalance(sms *SmsCloneCredential) (response string, err error) {
	preparedURL := util.SmsCloneCheckBalanceURL+"?username="+sms.Username+
	"&password="+sms.Password

	// contact endpoint
	var result []byte
   result, err = util.ContactEndpoint(preparedURL)
   if err != nil {
   	return response, err
   }

   response = string(result)
   return response, nil
}

func (*SmsCloneRepository) ValidateSmsCloneInput(smsInfo *SmsCloneNotification) (err map[string]interface{}) {
	err = make(map[string]interface{})

	if smsInfo.Username == "" {
		err["Username"] = util.ErrMissingUsername
	}

	if smsInfo.Sender == "" {
		err["Sender"] = util.ErrMissingSender
	}

	if smsInfo.Recipient == "" {
		err["Recipient"] = util.ErrMissingRecipient
	}

	if smsInfo.Message == "" {
		err["Message"] = util.ErrMissingMessage
	}

	if smsInfo.Password == "" {
		err["Password"] = util.ErrMissingPassword
	}

	return
}

func (*SmsCloneRepository) ValidateSmsCloneCredentials(smsInfo *SmsCloneCredential) (err map[string]interface{}) {
	err = make(map[string]interface{})

	if smsInfo.Username == "" {
		err["Username"] = util.ErrMissingUsername
	}

	if smsInfo.Password == "" {
		err["Password"] = util.ErrMissingPassword
	}

	return
}

func ResponseParser(text string) (parsedResponse SmsCloneResponse, err error) {
	result := strings.Split(text, "|")

	if len(result) <= 1 {
		err = errors.New(util.PossibleCredentialsErr)
		return SmsCloneResponse{}, err
	}

	first := strings.Split(result[0], "-")
	second := strings.Split(result[0], ":")
	third := strings.Split(first[1], ":")

	parsedResponse.BatchCode = first[0]
	parsedResponse.BatchDescription = third[0]
	parsedResponse.StatusCode = second[1]

	parsedResponse.Recipient = result[1]
	parsedResponse.MessageID = result[2]
	parsedResponse.MessageStatus = result[3]
	parsedResponse.StatusDescription = result[4]

	return
}


