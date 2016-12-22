package bean

import ()

// ProfileSearchCondition user's profile search struct
type ProfileSearchCondition struct {
	// select key     -1 is all
	SerialNumber   int `json:"serialNumber"`   // /global/bean/category.go Category/SubCategory SerialNumber
	HoursBilled    int `json:"hoursBilled"`    // /gather/global/business_const.go : Hours Billed
	AvailableHours int `json:"availableHours"` // /gather/global/business_const.go : Available Hours
	JobSuccess     int `json:"jobSuccess"`     // /gather/global/business_const.go : Job Success
	LastActivity   int `json:"lastActivity"`   // /gather/global/business_const.go : Last Activity
	FreelancerType int `json:"freelancerType"` // /gather/global/business_const.go : Freelancer Type
	HourlyRate     int `json:"hourlyRate"`     // /gather/global/business_const.go : Hours Rate
	RegionID       int `json:"regionID"`       // /gather/global/base_const.go : Region ID

	SearchKey string `json:"searchKey"` // searchkey
}
