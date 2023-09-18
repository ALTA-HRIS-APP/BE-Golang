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

// UpdateKaryawan implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) UpdateKaryawan(input reimbusment.ReimbursementEntity, id string) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&Reimbursement{}).Where("id=? and user_id=?", id,input.UserID).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("update data reimbursment error, hanya boleh mengedit reimbursment sendiri")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// SelectById implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) SelectById(id string) (int, error) {
	var inputModel Reimbursement
	tx := repo.db.Where("id=?",id).First(&inputModel)
	if tx.Error != nil {
		return 0, errors.New("error get batasan reimbursment")
	}
	return inputModel.BatasanReimburs, nil
}

// Update implements reimbusment.ReimbusmentDataInterface.
func (repo *ReimbusmentData) Update(input reimbusment.ReimbursementEntity, id string) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&Reimbursement{}).Where("id=?", id).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("update data reimbursment")
	}
	if tx.RowsAffected == 0 {
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
