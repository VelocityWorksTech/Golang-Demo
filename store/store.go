package store

import "github.com/jinzhu/gorm"

//OpenDBconnection To open & setup db connection
func OpenDBconnection() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "velocityworks.db")
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Catalog{}, &Distribution{}, &Publisher{},
		&ContactPoint{}, &Dataset{}).Error
	if err != nil {
		return nil, err
	}
	return db, nil
}
