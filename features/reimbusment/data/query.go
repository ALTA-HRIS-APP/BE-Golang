package data

import (
	"be_golang/klp3/features/reimbusment"
	"be_golang/klp3/helper"
	"errors"

	"gorm.io/gorm"
)

type ReimbusmentData struct {
	db *gorm.DB
}

// Update implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) Update(input reimbusment.ReimbursementEntity, id string) error {
	inputModel:=EntityToModel(input)
	tx:=repo.db.Model(&Reimbursement{}).Where("id=?",id).Updates(inputModel)
	if tx.Error != nil{
		return errors.New("update data reimbursment")
	}
	if tx.RowsAffected ==0{
		return errors.New("row not affected")
	}
	return nil
}

// Insert implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) Insert(input reimbusment.ReimbursementEntity) error {
	idUUID, errUUID := helper.GenerateUUID()
	if errUUID != nil {
		return errors.New("failed generated uuid")
	}
	inputModel := EntityToModel(input)
	inputModel.ID = idUUID
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

func New(db *gorm.DB) reimbusment.ReimbusmentDataInterface {
	return &ReimbusmentData{
		db: db,
	}
}
