package model

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/buger/jsonparser"
	"gorm.io/gorm"
)

// Catalog - first level in the payload hierarchy.
type Catalog struct {
	gorm.Model
	ConformsTo   string    `json:"conformsTo,omitempty"`
	DescribedBy  string    `json:"describedBy,omitempty"`
	Context      string    `json:"@context,omitempty"`
	MetadataType string    `json:"@type,omitempty"`
	Dataset      []Dataset `json:"dataset,omitempty"`
}

// Dataset stores each entry in dataset field of catalog.
type Dataset struct {
	gorm.Model
	CatalogID     uint           `json:"-"`
	MetadataType  string         `json:"@type,omitempty"`
	Title         string         `json:"title,omitempty"`
	Description   string         `json:"description,omitempty"`
	Modified      string         `json:"modified,omitempty"`
	AccessLevel   string         `json:"accessLevel,omitempty"`
	Identifier    string         `json:"identifier,omitempty"`
	License       string         `json:"license,omitempty"`
	Publisher     Publisher      `json:"publisher,omitempty"`
	ContactPoint  ContactPoint   `json:"contactPoint,omitempty"`
	Distributions []Distribution `json:"distribution,omitempty"`
	Keywords      string         `json:"keyword"`
	BureauCodes   string         `json:"bureauCode"`
	ProgramCodes  string         `json:"programCode"`
}

// Publisher will store all the information on the publisher field.
type Publisher struct {
	gorm.Model
	DatasetID    uint   `json:"-"`
	MetaDataType string `json:"@type,omitempty"`
	Name         string `json:"name,omitempty"`
	//	SubOrganizationOf string `json:"subOrganizationOf,omitempty"`
}

// ContactPoint will store all the information on the contactPoint field.
type ContactPoint struct {
	gorm.Model
	DatasetID    uint   `json:"-"`
	MetaDataType string `json:"@type,omitempty"`
	Fn           string `json:"fn,omitempty"`
	HasEmail     string `json:"hasEmail,omitempty"`
}

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

// ParseCatalogResponse parse json response.
func ParseCatalogResponse(data []byte, me *Catalog) error {

	var err error

	if me.MetadataType, err = jsonparser.GetString(data, "@type"); err != nil {
		return err
	}

	me.ConformsTo, _ = jsonparser.GetString(data, "conformsTo")
	me.DescribedBy, _ = jsonparser.GetString(data, "describedBy")
	me.Context, _ = jsonparser.GetString(data, "@context")

	datasets, dataType, _, _ := jsonparser.Get(data, "dataset")

	if dataType == jsonparser.Array {
		_, err = jsonparser.ArrayEach(
			datasets,
			func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				dataset := &Dataset{}
				if err = ParseDatasetResponse(value, dataset); err != nil {
					return
				}
				if dataset != nil {
					me.Dataset = append(me.Dataset, *dataset)
				}
			})
		if err != nil {
			return err
		}
	}

	return nil
}

//ParseDatasetResponse parse dataset
func ParseDatasetResponse(data []byte, me *Dataset) error {
	var err error
	me.MetadataType, err = jsonparser.GetString(data, "@type")
	if err != nil {
		return err
	}
	me.Title, _ = jsonparser.GetString(data, "title")
	me.Description, _ = jsonparser.GetString(data, "description")
	me.Modified, _ = jsonparser.GetString(data, "modified")
	me.AccessLevel, _ = jsonparser.GetString(data, "accessLevel")
	me.Identifier, _ = jsonparser.GetString(data, "identifier")
	me.License, _ = jsonparser.GetString(data, "license")

	//publisher
	me.Publisher = Publisher{}
	me.Publisher.MetaDataType, _ = jsonparser.GetString(data, "publisher", "@type")
	me.Publisher.Name, _ = jsonparser.GetString(data, "publisher", "name")

	//contact point
	me.ContactPoint = ContactPoint{}
	me.ContactPoint.MetaDataType, _ = jsonparser.GetString(data, "contactPoint", "@type")
	me.ContactPoint.Fn, _ = jsonparser.GetString(data, "contactPoint", "fn")
	me.ContactPoint.HasEmail, _ = jsonparser.GetString(data, "contactPoint", "hasEmail")

	me.Distributions = []Distribution{}
	if distributionsData, _, _, err := jsonparser.Get(data, "distribution"); err == nil {
		json.NewDecoder(bytes.NewBuffer(distributionsData)).Decode(&me.Distributions)
	}

	var strArray []string
	keywordsData, _, _, _ := jsonparser.Get(data, "keyword")
	json.NewDecoder(bytes.NewBuffer(keywordsData)).Decode(&strArray)
	me.Keywords = strings.Join(strArray, ",")

	bureauCodeData, _, _, _ := jsonparser.Get(data, "bureauCode")
	json.NewDecoder(bytes.NewBuffer(bureauCodeData)).Decode(&strArray)
	me.BureauCodes = strings.Join(strArray, ",")

	programCodesData, _, _, _ := jsonparser.Get(data, "programCode")
	json.NewDecoder(bytes.NewBuffer(programCodesData)).Decode(&strArray)
	me.ProgramCodes = strings.Join(strArray, ",")
	return err
}
