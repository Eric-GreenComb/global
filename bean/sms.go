package bean

import ()

// SMS alidayu sms struct
type SMS struct {
	RecNum          string `json:"recNum"`
	SmsFreeSignName string `json:"smsFreeSignName"`
	SmsTemplateCode string `json:"smsTemplateCode"`
	SmsParam        string `json:"smsParam"`
}
