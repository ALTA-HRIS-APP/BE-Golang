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

// Insert implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) Insert(input reimbusment.ReimbursementEntity) (string, error) {
	idUUID,errUUID:=helper.GenerateUUID()
	if errUUID != nil{
		return "",errors.New("failed generated uuid")
	}
	inputModel:=EntityToModel(input)
	inputModel.ID = idUUID
	tx:=repo.db.Create(&inputModel)
	if tx.Error !=nil{
		return "",errors.New("create failed")
	}
	return inputModel.ID,nil
}

// SelectUser implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) SelectUser(UserID string) (reimbusment.UserEntity, error) {
	var input User
	tx:=repo.db.Where("id=?",UserID).First(&input)
	if tx.Error != nil{
		return reimbusment.UserEntity{},tx.Error
	}
	output:=UserModelToEntity(input)
	return output,nil
}

func New(db *gorm.DB) reimbusment.ReimbusmentDataInterface {
	return &ReimbusmentData{
		db: db,
	}
}
