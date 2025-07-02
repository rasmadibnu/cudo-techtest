package repository

import (
	"cudo-techtest/config"
	"cudo-techtest/entity"
	"net/http"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	config config.Database
}

func NewTransactionRepository(database config.Database) TransactionRepository {
	return TransactionRepository{
		config: database,
	}
}

// @Summary : Define model
// @Description :
// @Author : rasmadibnu
func (r *TransactionRepository) Model(req *http.Request) *gorm.DB {
	return r.config.DB.Model(&entity.Transaction{}).Scopes(r.Filter(req)).Preload("PreloadName").Order("column desc")
}

// @Summary : Find Data
// @Description :
// @Author : rasmadibnu
func (r *TransactionRepository) Find() ([]entity.Transaction, error) {
	var tx []entity.Transaction

	err := r.config.DB.Find(&tx).Error

	if err != nil {
		return tx, err
	}

	return tx, nil
}

// @Summary : Define Filter
// @Description : Define Filter
// @Author : rasmadibnu
func (s *TransactionRepository) Filter(r *http.Request) func(db *gorm.DB) *gorm.DB {
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
func (r *TransactionRepository) Insert(ety entity.Transaction) (entity.Transaction, error) {
	err := r.config.DB.Debug().Create(&ety).Error

	if err != nil {
		return ety, err
	}

	return ety, nil
}

// @Summary : Find
// @Description : Return single data by pk
// @Author : rasmadibnu
func (r *TransactionRepository) FindById(ID int) (entity.Transaction, error) {
	var ety entity.Transaction
	err := r.config.DB.Where(ID).First(&ety).Error

	if err != nil {
		return ety, err
	}

	return ety, nil
}

// @Summary : Update
// @Description : Update data by pk
// @Author : rasmadibnu
func (r *TransactionRepository) Update(ety entity.Transaction, ID int) (entity.Transaction, error) {
	err := r.config.DB.Debug().Where(ID).Updates(&ety).Error

	if err != nil {
		return ety, err
	}

	return ety, nil
}

// @Summary : Delete
// @Description : Delete data by pk
// @Author : rasmadibnu
func (r *TransactionRepository) Delete(ID int) (bool, error) {
	var ety entity.Transaction

	err := r.config.DB.Debug().Where(ID).Delete(&ety).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
