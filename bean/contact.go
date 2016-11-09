package bean

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Contact user-client contact
type Contact struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	ClientEmail     string        `form:"client_email" bson:"client_email" json:"client_email"`
	FreeLancerEmail string        `form:"freelancer_email" bson:"freelancer_email" json:"freelancer_email"`

	ContactContent string `bson:"contact_content" json:"contact_content"`
	ContactTpl     string `form:"contact_tpl" bson:"contact_tpl" json:"contact_tpl"`
	TplParam       string `form:"tpl_param" bson:"tpl_param" json:"tpl_param"`

	ClientSignup        bool   `bson:"client_signup" json:"client_signup"`
	ClientSignature     string `bson:"client_signature" json:"client_signature"`
	FreelancerSignup    bool   `bson:"freelancer_signup" json:"freelancer_signup"`
	FreelancerSignature string `bson:"freelancer_signature" json:"freelancer_signature"`
	Dealed              bool   `bson:"dealed" json:"dealed"`

	DealedTime  time.Time `bson:"dealedtime" json:"dealedtime"`
	CreatedTime time.Time `bson:"createdtime" json:"createdtime"`
}
