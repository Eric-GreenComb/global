package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// User user struct
type User struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`                         // 必须要设置bson:"_id" 不然mgo不会认为是主键
	Invited     bson.ObjectId `form:"invited" bson:"invited" json:"invited"` // invited by someone
	Email       string        `form:"email" bson:"email" json:"email"`       // 全是小写
	Pwd         string        `form:"pwd" bson:"pwd" json:"-"`
	Actived     bool          `bson:"actived" json:"actived"`
	CreatedTime int64         `bson:"createTime" json:"createTime"`
}

// UserDto user readonly struct
type UserDto struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`          // 必须要设置bson:"_id" 不然mgo不会认为是主键
	Invited bson.ObjectId `bson:"invited" json:"invited"` // invited by someone
	Email   string        `bson:"email" json:"email"`     // 全是小写
}
