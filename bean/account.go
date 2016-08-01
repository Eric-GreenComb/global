package bean

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Account struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserId bson.ObjectId `bson:"user_id" json:"user_id"`
	Email  string        `bson:"email" json:"email"`

	MultiCurrency []MultiCurrencyAccount `bson:"multi_curreency" json:"multi_curreency"`

	CreatedTime time.Time `bson:"createdtime" json:"createdtime"`
}

type MultiCurrencyAccount struct {
	Currency string `bson:"currency" json:"currency"`
	Amount   int64  `bson:"amount" json:"amount"`
}

type Billing struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserId    bson.ObjectId `bson:"user_id" json:"user_id"`
	PayUserId bson.ObjectId `bson:"pay_user_id" json:"pay_user_id"`
	ProfileId bson.ObjectId `bson:"profile_id" json:"profile_id"`
	Memo      string        `bson:"memo" json:"memo"`
	CostType  int           `bson:"cost_type" json:"cost_type"`

	Operate   int    `bson:"operate" json:"operate"` // 1 or -1 Operate
	Currency  string `bson:"currency" json:"currency"`
	Amount    int64  `bson:"amount" json:"amount"`
	PayType   int    `bson:"paytype" json:"paytype"`
	PayFee    int64  `bson:"payfee" json:"payfee"`
	RealCosts int64  `bson:"real_costs" json:"real_costs"`

	Status      int       `bson:"status" json:"status"`
	CreatedTime time.Time `bson:"createdtime" json:"createdtime"`
}
