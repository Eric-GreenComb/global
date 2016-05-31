package bean

import (
	"labix.org/v2/mgo/bson"
)

type WorkHistory struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	ProfileID bson.ObjectId `bson:"profile_id" json:"profile_id"`

	HistoryAndFeedbacks []WorkHistoryAndFeedback `bson:"history_feedbacks" json:"history_feedbacks"`
}

type WorkHistoryAndFeedback struct {
	Title      string `bson:"title" json:"title"`
	WorkPeriod string `bson:"work_period" json:"work_period"`

	WorkType   string `bson:"work_type" json:"work_type"` // Fixed Price or $27.78 / hr
	WorkEarned string `bson:"work_earned" json:"work_earned"`
	WorkHours  int    `bson:"work_hours" json:"work_hours"`

	WorkFeedbacks []WorkFeedback `bson:"work_feedbacks" json:"work_feedbacks"`
}

type WorkFeedback struct {
	WorkRate int    `bson:"workrate" json:"workrate"`
	Feedback string `bson:"feedback" json:"feedback"`
	Writeby  string `bson:"writeby" json:"writeby"`
}
