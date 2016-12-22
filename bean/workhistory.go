package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// WorkHistory user work history on banerwai
type WorkHistory struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	ProfileID bson.ObjectId `bson:"profileID" json:"profileID"`

	HistoryAndFeedbacks []WorkHistoryAndFeedback `bson:"historyFeedbacks" json:"historyFeedbacks"`
}

// WorkHistoryAndFeedback user work history / feedback
type WorkHistoryAndFeedback struct {
	Title      string `bson:"title" json:"title"`
	WorkPeriod string `bson:"workPeriod" json:"workPeriod"`

	WorkType   string `bson:"workType" json:"workType"` // Fixed Price or $27.78 / hr
	WorkEarned string `bson:"workEarned" json:"workEarned"`
	WorkHours  int    `bson:"workHours" json:"workHours"`

	WorkFeedbacks []WorkFeedback `bson:"workFeedbacks" json:"workFeedbacks"`
}

// WorkFeedback user work feedback
type WorkFeedback struct {
	WorkRate int    `bson:"workrate" json:"workrate"`
	Feedback string `bson:"feedback" json:"feedback"`
	Writeby  string `bson:"writeby" json:"writeby"`
}
