package data

import (
	"be_golang/klp3/features/cuti"
	"be_golang/klp3/helper"
	"errors"

	"gorm.io/gorm"
)

type CutiData struct {
	db *gorm.DB
}

// insert implements cuti.CutiDataInterface.
func (repo *CutiData) Insert(input cuti.CutiEntity) error {
	inputModel := EntityToModel(input)
	id, err := helper.GenerateUUID()
	if err != nil {
		return errors.New("error create uuid")
	}
	inputModel.ID = id
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return errors.New("create error")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affective")
	}
	return nil
}

func New(db *gorm.DB) cuti.CutiDataInterface {
	return &CutiData{
		db: db,
	}
}
