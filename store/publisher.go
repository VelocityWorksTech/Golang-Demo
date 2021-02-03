package store

import "github.com/jinzhu/gorm"

// Publisher will store all the information on the publisher field.
type Publisher struct {
	gorm.Model
	DatasetID    uint   `json:"-"`
	MetaDataType string `json:"@type,omitempty"`
	Name         string `json:"name,omitempty"`
}
