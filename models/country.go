package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Country struct {
	Model

	ID_str      string         `json:"id_str"`
	Alpha2Code  string         `json:"alpha_2_code"`
	Alpha3Code  string         `json:"alpha_3_code"`
	Name        string         `json:"name"`
	Demonym     sql.NullString `json:"demonym"`
	ContinentId string         `json:"continent_ID"`
}

// ExistCountryByID checks if an country exists based on ID
func ExistCountryByID(id int) (bool, error) {
	var country Country
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&country).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if country.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetCountryTotal gets the total number of countries based on the constraints
func GetCountryTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Country{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetCountrys gets a list of countries based on paging constraints
func GetCountrys(pageNum int, pageSize int, maps interface{}) ([]*Country, error) {
	var countries []*Country
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&countries).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return countries, nil
}

// GetCountry Get a single country based on ID
func GetCountry(id int) (*Country, error) {
	var country Country
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&country).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&country).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &country, nil
}

// EditCountry modify a single country
func EditCountry(id int, data interface{}) error {
	if err := db.Model(&Country{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddCountry add a single country
func AddCountry(data map[string]interface{}) error {

	country := Country{
		ID_str:      data["id_str"].(string),
		Alpha2Code:  data["alpha_2_code"].(string),
		Alpha3Code:  data["alpha_3_code"].(string),
		Name:        data["name"].(string),
		Demonym:     data["demonym"].(sql.NullString),
		ContinentId: data["continent_ID"].(string),
	}

	if err := db.Create(&country).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCountry delete a single country
func DeleteCountry(id int) error {
	if err := db.Where("id = ?", id).Delete(Country{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllCountry clear all country
func CleanAllCountry() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Country{}).Error; err != nil {
		return err
	}

	return nil
}
