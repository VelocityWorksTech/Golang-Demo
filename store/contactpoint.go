package store

import "github.com/jinzhu/gorm"

// ContactPoint will store all the information on the contactPoint field.
type ContactPoint struct {
	gorm.Model
	DatasetID    uint   `json:"-"`
	MetaDataType string `json:"@type,omitempty"`
	Fn           string `json:"fn,omitempty"`
	HasEmail     string `json:"hasEmail,omitempty"`
}
