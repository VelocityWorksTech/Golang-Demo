package store

import (
	"github.com/buger/jsonparser"
	"github.com/jinzhu/gorm"
)


// Catalog - first level in the payload hierarchy.
type Catalog struct {
	gorm.Model

	URL          string    `json:"-"`
	ConformsTo   string    `json:"conformsTo,omitempty"`
	DescribedBy  string    `json:"describedBy,omitempty"`
	Context      string    `json:"@context,omitempty"`
	MetadataType string    `json:"@type,omitempty"`
	Dataset      []Dataset `json:"dataset,omitempty"`
}


// Parse json response.
func (me *Catalog) Parse(data []byte) error {

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
				if err = dataset.Parse(value); err != nil {
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
