package service

import (
	"be_golang/klp3/features/reimbusment"
	"errors"

	"github.com/go-playground/validator/v10"
)

type ReimbursementService struct {
	reimbursmentService reimbusment.ReimbusmentDataInterface
	validate            *validator.Validate
}

// Add implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Add(input reimbusment.ReimbursementEntity) (error) {
	errValidate:=service.validate.Struct(input)
	if errValidate != nil{
		return errors.New("error validate, data deskripsi, nominal, tipe reimbusment required")
	}
	_,errUser:=service.reimbursmentService.SelectUser(input.UserID)
	if errUser != nil{
		return errors.New("user not found")
	}
	_,errInsert:=service.reimbursmentService.Insert(input)
	if errInsert != nil{
		return errInsert
	}
	return nil
}

func New(service reimbusment.ReimbusmentDataInterface) reimbusment.ReimbusmentServiceInterface {
	return &ReimbursementService{
		reimbursmentService: service,
		validate:            validator.New(),
	}
}
