package store

import "github.com/jinzhu/gorm"

// Distribution will store all the information on the distribution field.
type Distribution struct {
	gorm.Model
	DatasetID       uint   `json:"-"`
	MetaDataType    string `json:"@type,omitempty"`
	DownloadURL     string `json:"downloadURL,omitempty"`
	AccessURL       string `json:"accessURL,omitempty"`
	MediaType       string `json:"mediaType,omitempty"`
	Format          string `json:"format,omitempty"`
	Title           string `json:"title,omitempty"`
	Description     string `json:"description,omitempty"`
	DescribedBy     string `json:"describedBy,omitempty"`
	DescribedByType string `json:"describedByType,omitempty"`
	ConformsTo      string `json:"conformsTo,omitempty"`
}
