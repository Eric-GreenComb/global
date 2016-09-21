package bean

import ()

// Category profile category
type Category struct {
	ID            string        `json:"id"`
	SerialNumber  int32         `json:"serialnumber"`
	Name          string        `json:"name"`
	Desc          string        `json:"desc"`
	Fa            string        `json:"fa"`
	Index         bool          `json:"index"`
	Subcategories []SubCategory `json:"subcategories"`
}

// SubCategory profile subcategory
type SubCategory struct {
	ID           string `json:"id"`
	SerialNumber int32  `json:"serialnumber"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
}
