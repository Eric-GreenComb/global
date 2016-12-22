package bean

import (
	"gopkg.in/mgo.v2/bson"
)

// Profile user profile
type Profile struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID bson.ObjectId `bson:"userID" json:"userID"`

	// select key     -1 is all
	CategoryNumber int `bson:"categoryNumber" json:"categoryNumber"` // /global/bean/category.go Category/SubCategory SerialNumber
	SerialNumber   int `bson:"serialNumber" json:"serialNumber"`     // /global/bean/category.go Category/SubCategory SerialNumber

	HoursBilled     int `bson:"hourBilled" json:"hourBilled"`           // /global/business_const.go : Hours Billed
	AvailableHours  int `bson:"availableHours" json:"availableHours"`   // /global/business_const.go : Available Hours
	JobSuccess      int `bson:"jobSuccess" json:"jobSuccess"`           // /global/business_const.go : Job Success
	LastActivity    int `bson:"lastActivity" json:"lastActivity"`       // /global/business_const.go : Last Activity
	FreelancerType  int `bson:"freelancerType" json:"freelancerType"`   // /global/business_const.go : Freelancer Type
	HourlyRate      int `bson:"hourlyRate" json:"hourlyRate"`           // /global/business_const.go : Hours Rate
	RegionID        int `bson:"regionID" json:"regionID"`               // /global/base_const.go : Region ID
	ExperienceLevel int `bson:"experienceLevel" json:"experienceLevel"` // /global/business_const.go : Skill Experience Level

	// show
	Name              string         `bson:"freelancerName" json:"freelancerName"`
	JobTitle          string         `bson:"jobTitle" json:"jobTitle"`
	Overview          string         `bson:"overview" json:"overview"` // searchkey
	HourRate          int            `bson:"hourRate" json:"hourRate"`
	WorkHours         int            `bson:"workHours" json:"workHours"`
	PortfolioNums     int            `bson:"portfolioNums" json:"portfolioNums"`
	SkillTags         string         `bson:"skillTags" json:"skillTags"`
	AgencyHourlyRates string         `bson:"agencyHourlyRates" json:"agencyHourlyRates"`
	AgencyMembers     []AgencyMember `bson:"agencyMembers" json:"agencyMembers"`

	Status         bool  `bson:"status" json:"status"`
	CreatedTime    int64 `bson:"createTime" json:"createTime"`
	LastActiveTime int64 `bson:"lastActiveTime" json:"lastActiveTime"`
}

// AgencyMember if user as a client ,client's member
type AgencyMember struct {
	Names     string `bson:"profileName" json:"profileName"`
	JobTitles string `bson:"profileJobTitle" json:"profileJobTitle"`
	Manager   bool   `bson:"manager" json:"manager"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" json:"phone"`
}

// ProfileDto profile read struct
type ProfileDto struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID bson.ObjectId `bson:"userID" json:"userID"`

	Name              string         `bson:"freelancerName" json:"freelancerName"`
	JobTitle          string         `bson:"jobTitle" json:"jobTitle"`
	Overview          string         `bson:"overview" json:"overview"` // searchkey
	WorkHours         int            `bson:"workHours" json:"workHours"`
	PortfolioNums     int            `bson:"portfolioNums" json:"portfolioNums"`
	SkillTags         string         `bson:"skillTags" json:"skillTags"`
	AgencyHourlyRates string         `bson:"agencyHourlyRates" json:"agencyHourlyRates"`
	AgencyMembers     []AgencyMember `bson:"agencyMembers" json:"agencyMembers"`

	CreatedTime int64 `bson:"createTime" json:"createTime"`
}

// ProfileSearchDto profile search dto
type ProfileSearchDto struct {
	SerialNumber   int    `json:"serialNumber"`
	HoursBilled    int    `json:"hoursBilled"`
	AvailableHours int    `json:"availableHours"`
	JobSuccess     int    `json:"jobSuccess"`
	LastActivity   int    `json:"lastAtivity"`
	FreelancerType int    `json:"freelancerType"`
	HourlyRate     int    `json:"hourlyRate"`
	RegionID       int    `json:"regionID"`
	SearchKey      string `json:"searchKey"`
}
