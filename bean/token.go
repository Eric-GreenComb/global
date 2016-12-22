package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// 随机token

// token type
const (
	TokenPwd = iota // 0
	TokenActiveEmail
	TokenUpdateEmail
)

// 过期时间 hour
const (
	PwdOverHours         = 2 * 60 * 60
	ActiveEmailOverHours = 48 * 60 * 60
	UpdateEmailOverHours = 2 * 60 * 60
)

// Token token struct
type Token struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Key         string        `bson:"key" json:"key"`
	Token       string        `bson:"token" json:"token"`
	Type        int64         `bson:"type" json:"type"`
	CreatedTime int64         `bson:"createTime" json:"createTime"`
}
