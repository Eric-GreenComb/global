package bean

import ()

// APIRequest API Request
// Header : request header
// Payload : request payload
// Signature : request signature
type APIRequest struct {
	Header    APIRequestHeader    `json:"header,omitempty"`
	Payload   APIRequestPayload   `json:"payload,omitempty"`
	Signature APIRequestSignature `json:"signature,omitempty"`
}

// APIRequestHeader API request header
// Method : invoke, query ...
// Ver : api version 1.0.0
type APIRequestHeader struct {
	Method string `json:"method,omitempty"`
	Ver    string `json:"ver,omitempty"`
}

// APIRequestPayload API request payload
// Iss : 签发者
// Iat : 在什么时候签发的
// Aud : 接收的一方
// Params : form params map
// SortKeys : map sort keys
type APIRequestPayload struct {
	Iss    string            `json:"iss,omitempty"`
	Iat    int64             `json:"iat,omitempty"`
	Aud    string            `json:"aud,omitempty"`
	Params map[string]string `json:"params,omitempty"`
}

// APIRequestSignature API request signature
// Sign iat + . + sort params
// Alg  md5/
type APIRequestSignature struct {
	Sign string `json:"sign,omitempty"`
	Alg  string `json:"alg,omitempty"`
}
