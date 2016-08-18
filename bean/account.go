package bean

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Account struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserId bson.ObjectId `form:"user_id" bson:"user_id" json:"user_id"`
	Email  string        `form:"email" bson:"email" json:"email"`

	MultiCurrency []MultiCurrencyAccount `bson:"multi_curreency" json:"multi_curreency"`

	CreatedTime time.Time `bson:"createdtime" json:"createdtime"`
}

type MultiCurrencyAccount struct {
	Currency string `bson:"currency" json:"currency"`
	Amount   int64  `bson:"amount" json:"amount"`
}

type Billing struct {
	Id        bson.ObjectId `form:"id" bson:"_id,omitempty" json:"_id"`
	UserId    bson.ObjectId `form:"user_id" bson:"user_id" json:"user_id"`
	PayUserId bson.ObjectId `form:"pay_user_id" bson:"pay_user_id" json:"pay_user_id"`
	ServiceId bson.ObjectId `form:"service_id" bson:"service_id" json:"service_id"`
	CostType  int           `form:"cost_type" bson:"cost_type" json:"cost_type"`
	Memo      string        `form:"memo" bson:"memo" json:"memo"`

	Operate     int    `bson:"operate" json:"operate"` // 1 or -1 Operate
	Currency    string `form:"currency" bson:"currency" json:"currency"`
	Amount      int64  `form:"amount" bson:"amount" json:"amount"`
	PayType     int    `form:"paytype" bson:"paytype" json:"paytype"`
	PayFee      int64  `bson:"payfee" json:"payfee"`
	PlatformFee int64  `bson:"platform_fee" json:"platform_fee"`
	RealAmount  int64  `bson:"real_amount" json:"real_amount"`

	Status      int       `bson:"status" json:"status"`
	CreatedTime time.Time `bson:"createdtime" json:"createdtime"`
}
