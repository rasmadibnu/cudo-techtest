package repository

import (
	"cudo-techtest/config"
	"cudo-techtest/entity"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository struct {
	config config.Database
}

func NewUserRepository(database config.Database) UserRepository {
	return UserRepository{
		config: database,
	}
}

// @Summary : Define model
// @Description :
// @Author : rasmadibnu
func (r *UserRepository) Model(req *http.Request) *gorm.DB {
	return r.config.DB.Model(&entity.User{}).Scopes(r.Filter(req)).Preload("PreloadName").Order("column desc")
}

// @Summary : Define Filter
// @Description : Define Filter
// @Author : rasmadibnu
func (s *UserRepository) Filter(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()

		if q.Has("from") && q.Has("to") {
			db = db.Where("column between ? and ?", q.Get("from"), q.Get("to"))
		}

		if q.Has("param") {

			db = db.Where("column in (?)", q["param"])
		}

		if q.Has("param") {

			db = db.Where("column in (?)", q["param"])
		}

		if q.Has("param") {

			db = db.Where("column like '%" + q.Get("column") + "%'")
		}

		return db.Debug()
	}
}

// @Summary : Insert
// @Description : Insert to database
// @Author : rasmadibnu
func (r *UserRepository) Insert(ety entity.User) (entity.User, error) {
	err := r.config.DB.Debug().Create(&ety).Error

	if err != nil {
		return ety, err
	}

	return ety, nil
}

// @Summary : Find
// @Description : Return single data by pk
// @Author : rasmadibnu
func (r *UserRepository) FindById(ID int) (entity.User, error) {
	var ety entity.User
	err := r.config.DB.Where(ID).First(&ety).Error

	if err != nil {
		return ety, err
	}

	return ety, nil
}

// @Summary : Update
// @Description : Update data by pk
// @Author : rasmadibnu
func (r *UserRepository) Update(ety entity.User, ID int) (entity.User, error) {
	err := r.config.DB.Debug().Where(ID).Updates(&ety).Error

	if err != nil {
		return ety, err
	}

	return ety, nil
}

// @Summary : Delete
// @Description : Delete data by pk
// @Author : rasmadibnu
func (r *UserRepository) Delete(ID int) (bool, error) {
	var ety entity.User

	err := r.config.DB.Debug().Where(ID).Delete(&ety).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
