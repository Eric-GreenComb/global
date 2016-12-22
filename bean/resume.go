package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// Resume user's resume struct
type Resume struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID    bson.ObjectId `bson:"userID" json:"userID"`
	AuthEmail string        `bson:"authEmail" json:"authEmail"`
	Phone     string        `bson:"phone" json:"phone"`

	SkillExperiences    []SkillExperience   `bson:"experienceLevels" json:"experienceLevels"`
	ToolandArchs        []ToolandArch       `bson:"toolArchs" json:"toolArchs"`
	Portfolioes         []Portfolio         `bson:"portfolioes" json:"portfolioes"`
	EmploymentHistories []EmploymentHistory `bson:"employmentHistories" json:"employmentHistories"`
	Educations          []Education         `bson:"educations" json:"educations"`
	OtherExperiences    []OtherExperience   `bson:"otherExperiences" json:"otherExperiences"`
}

// SkillExperience user skill experience
type SkillExperience struct {
	SkillTitle string `bson:"skillTitle" json:"skillTitle"`
	SkillLevel int    `bson:"skillLevel" json:"skillLevel"` // years
}

// ToolandArch user use tools / archs
type ToolandArch struct {
	ToolTitle string `bson:"tool_title" json:"tool_title"`
	ToolLevel int    `bson:"tool_level" json:"tool_level"` // years
}

// Portfolio user project portfolio
type Portfolio struct {
	ProjectTitle  string `bson:"projectTitle" json:"projectTitle"`
	Overview      string `bson:"overview" json:"overview"`
	CategoryID    string `bson:"categoryID" json:"categoryID"`
	SubCategoryID string `bson:"subcategoryID" json:"subcategoryID"`
	ProjectURL    string `bson:"projectURL" json:"projectURL"`
	SkillTags     string `bson:"skillTags" json:"skillTags"`
}

// EmploymentHistory user employment history
type EmploymentHistory struct {
	AgencyName  string `bson:"agencyName" json:"agencyName"`
	Location    string `bson:"locationCity" json:"locationCity"`
	WorkTitle   string `bson:"workTitle" json:"workTitle"`
	WorkPeriod  string `bson:"workPeriod" json:"workPeriod"`
	Description string `bson:"description" json:"description"`
}

// Education user education history
type Education struct {
	SchoolName  string `bson:"schoolName" json:"schoolName"`
	StudyPeriod string `bson:"studyPeriod" json:"studyPeriod"`
	Degree      string `bson:"degree" json:"degree"`
	StudyArea   string `bson:"studyArea" json:"studyArea"`
	Description string `bson:"description" json:"description"`
}

// OtherExperience user other experience
type OtherExperience struct {
	Subject     string `bson:"subject" json:"subject"`
	Description string `bson:"description" json:"description"`
}
