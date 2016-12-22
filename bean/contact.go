package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// Contact user-client contact
type Contact struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	ClientEmail     string        `form:"clientEmail" bson:"clientEmail" json:"clientEmail"`
	FreeLancerEmail string        `form:"freelancerEmail" bson:"freelancerEmail" json:"freelancerEmail"`

	ContactContent string `bson:"contactContent" json:"contactContent"`
	ContactTpl     string `form:"contactTpl" bson:"contactTpl" json:"contactTpl"`
	TplParam       string `form:"tplParam" bson:"tplParam" json:"tplParam"`

	ClientSignup        bool   `bson:"clientSignup" json:"clientSignup"`
	ClientSignature     string `bson:"clientSignature" json:"clientSignature"`
	FreelancerSignup    bool   `bson:"freelancerSignup" json:"freelancerSignup"`
	FreelancerSignature string `bson:"freelancerSignature" json:"freelancerSignature"`
	Dealed              bool   `bson:"dealed" json:"dealed"`

	DealedTime  int64 `bson:"dealTime" json:"dealTime"`
	CreatedTime int64 `bson:"createTime" json:"createTime"`
}
