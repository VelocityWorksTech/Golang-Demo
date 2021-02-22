package store

import (
	"github.com/buger/jsonparser"
	"github.com/jinzhu/gorm"
)

// Catalog - first level in the payload hierarchy.
type Catalog struct {
	gorm.Model

	URL          string    `json:"url,omitempty"`
	ConformsTo   string    `json:"conformsTo,omitempty"`
	DescribedBy  string    `json:"describedBy,omitempty"`
	Context      string    `json:"context,omitempty"`
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

//First Get first record
func (me *Catalog) First(db *gorm.DB) {
	db.First(me)
	db.Where("catalog_id=?", me.ID).Find(&me.Dataset)
	for i := 0; i < len(me.Dataset); i++ {
		db.Where("dataset_id=?", me.Dataset[i].ID).Find(&me.Dataset[i].Publisher)
		db.Where("dataset_id=?", me.Dataset[i].ID).Find(&me.Dataset[i].ContactPoint)
		db.Where("dataset_id=?", me.Dataset[i].ID).Find(&me.Dataset[i].Distributions)
	}
}

//Delete deletes the catalog and its hierarcy
func (me *Catalog) Delete(db *gorm.DB) {
	datasets := []int64{}
	db.Model(&Dataset{}).Where(&Dataset{}, "CatalogID = ?", me.ID).Pluck("ID", &datasets)
	for _, d := range datasets {
		db.Delete(&Publisher{}).Where("DatasetID=?", d)
		db.Delete(&Distribution{}).Where("DatasetID=?", d)
		db.Delete(&ContactPoint{}).Where("DatasetID=?", d)
	}
	db.Delete(&Dataset{}).Where("CatalogID=?", me.ID)
	db.Delete(me)
}

//GetDatasets retreives all datasets
func (me *Catalog) GetDatasets(db *gorm.DB) {
	db.Table("datasets").Where("catalog_id = ?", me.ID).Find(&me.Dataset)
	for i := 0; i < len(me.Dataset); i++ {
		db.Table("contact_points").Where("dataset_id = ?", me.Dataset[i].ID).Find(&me.Dataset[i].ContactPoint)
		db.Table("publishers").Where("dataset_id = ?", me.Dataset[i].ID).Find(&me.Dataset[i].Publisher)
		db.Table("distributions").Where("dataset_id = ?", me.Dataset[i].ID).Find(&me.Dataset[i].Distributions)
	}
}
