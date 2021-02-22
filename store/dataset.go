package store

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/jinzhu/gorm"
)

// Dataset stores each entry in dataset field of catalog.
type Dataset struct {
	gorm.Model
	CatalogID     uint           `json:"-"`
	MetadataType  string         `json:"type,omitempty"`
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

//Parse parse dataset
//Utility function
func (me *Dataset) Parse(data []byte) error {
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
