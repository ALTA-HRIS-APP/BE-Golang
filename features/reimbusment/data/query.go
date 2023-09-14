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

// UpdateStatusByHR implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) UpdateStatusByHR(status string, idUser string,id string) error {
	var inputModel Reimbursement
	inputModel.Persetujuan=status
	tx:=repo.db.Model(&Reimbursement{}).Where("user_id=? and id=?",idUser,id).Updates(inputModel)
	if tx.Error != nil{
		return errors.New("error status update by HR")
	}
	if tx.RowsAffected == 0{
		return errors.New("row not affected")
	}
	return nil
}

// UpdateStatusByManager implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) UpdateStatusByManager(status string, idUser string,id string) error {
	var inputModel Reimbursement
	inputModel.Status=status
	tx:=repo.db.Model(&Reimbursement{}).Where("user_id=? and id=?",idUser,id).Updates(inputModel)
	if tx.Error != nil{
		return errors.New("error status update by Manager")
	}
	if tx.RowsAffected == 0{
		return errors.New("row not affected")
	}
	return nil
}

// UpdateUser implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) UpdateUser(input reimbusment.ReimbursementEntity, idUser string,id string) error {
	inputModel:=EntityToModel(input)
	tx:=repo.db.Model(&Reimbursement{}).Where("user_id=? and id=?",idUser,id).Updates(inputModel)
	if tx.Error != nil{
		return errors.New("error data update by user karyawan")
	}
	if tx.RowsAffected == 0{
		return errors.New("row not affected")
	}
	return nil
}

// Insert implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) Insert(input reimbusment.ReimbursementEntity) (string, error) {
	idUUID, errUUID := helper.GenerateUUID()
	if errUUID != nil {
		return "", errors.New("failed generated uuid")
	}
	inputModel := EntityToModel(input)
	inputModel.ID = idUUID
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return "", errors.New("create failed")
	}
	if tx.RowsAffected == 0{
		return "",errors.New("row not affected")
	}
	return inputModel.ID, nil
}

// SelectUser implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) SelectUser(UserID string) (reimbusment.UserEntity, error) {
	var input User
	tx := repo.db.Where("id=?", UserID).First(&input)
	if tx.Error != nil {
		return reimbusment.UserEntity{}, tx.Error
	}
	if tx.RowsAffected == 0{
		return reimbusment.UserEntity{},errors.New("row not affected")
	}
	output := UserModelToEntity(input)
	return output, nil
}

func New(db *gorm.DB) reimbusment.ReimbusmentDataInterface {
	return &ReimbusmentData{
		db: db,
	}
}
