package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// Account account struct
type Account struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID bson.ObjectId `form:"userID" bson:"userID" json:"userID"`
	Email  string        `form:"email" bson:"email" json:"email"`

	MultiCurrency []MultiCurrencyAccount `bson:"multi_curreency" json:"multi_curreency"`

	CreatedTime int64 `bson:"createdtime" json:"createdtime"`
}

// MultiCurrencyAccount multi currency of account
type MultiCurrencyAccount struct {
	Currency string `bson:"currency" json:"currency"`
	Amount   int64  `bson:"amount" json:"amount"`
}

// Billing user account billing
type Billing struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID    bson.ObjectId `form:"userID" bson:"userID" json:"userID"`
	PayUserID bson.ObjectId `form:"payUserID" bson:"payUserID" json:"payUserID"`
	ServiceID bson.ObjectId `form:"serviceID" bson:"serviceID" json:"serviceID"`
	LinkID    bson.ObjectId `form:"linkID" bson:"linkID" json:"linkID"`
	CostType  int           `form:"costType" bson:"costType" json:"costType"`
	Memo      string        `form:"memo" bson:"memo" json:"memo"`

	Operate     int    `bson:"operate" json:"operate"` // 1 or -1 Operate
	Currency    string `form:"currency" bson:"currency" json:"currency"`
	Amount      int64  `form:"amount" bson:"amount" json:"amount"`
	PayType     int    `form:"payType" bson:"payType" json:"payType"`
	PayFee      int64  `bson:"payFee" json:"payFee"`
	PlatformFee int64  `bson:"platformFee" json:"platformFee"`
	RealAmount  int64  `bson:"realAmount" json:"realAmount"`

	Status      int   `bson:"status" json:"status"`
	CreatedTime int64 `bson:"createTime" json:"createTime"`
}
