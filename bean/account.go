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
	ProfileId bson.ObjectId `bson:"profile_id" json:"profile_id"`

	Operate   int    `bson:"operate" json:"operate"` // 1 or -1 Operate
	Currency  string `bson:"currency" json:"currency"`
	Amount    int64  `bson:"amount" json:"amount"`
	PayType   int    `bson:"paytype" json:"paytype"`
	PayFee    int    `bson:"payfee" json:"payfee"`
	RealCosts int    `bson:"real_costs" json:"real_costs"`

	Status      int       `bson:"status" json:"status"`
	CreatedTime time.Time `bson:"createdtime" json:"createdtime"`
}
