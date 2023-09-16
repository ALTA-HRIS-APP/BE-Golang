package service

import (
	"be_golang/klp3/features/reimbusment"
	usernodejs "be_golang/klp3/features/userNodejs"

	"errors"

	"github.com/go-playground/validator/v10"
)

type ReimbursementService struct {
	reimbursmentService reimbusment.ReimbusmentDataInterface
	validate            *validator.Validate
}

// Edit implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Edit(input reimbusment.ReimbursementEntity, id string) error {
	dataUser,errUser:=usernodejs.GetByIdUser(input.ID)
	if errUser !=nil{
		return errors.New("failed get user by id")
	}
	if dataUser.Jabatan =="karyawan"{
		if input.Persetujuan != ""{
			return errors.New("hanya hr yang bisa approve final")
		}
		if input.Status != ""{
			return errors.New("hanya manager yang bisa approve")
		}
		err:=service.reimbursmentService.Update(input,id)
		if err != nil{
			return err
		}
		return nil
	}else if dataUser.Jabatan == "manager"{
		if input.Persetujuan != ""{
			return errors.New("hanya hr yang bisa approve final")
		}
		if input.UserID == ""{
			return errors.New("harap pilih user id yang ingin di approve")
		}
		err:=service.reimbursmentService.Update(input,id)
		if err != nil{
			return err
		}
		return nil
	}else{
		if input.UserID == ""{
			return errors.New("harap pilih user id yang ingin di approve")
		}
		err:=service.reimbursmentService.Update(input,id)
		if err != nil{
			return err
		}
		return nil
	}
}

// Add implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Add(input reimbusment.ReimbursementEntity) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate, data deskripsi, nominal, tipe reimbusment required")
	}
	if input.Nominal > 5000000{
		return errors.New("pengajuan reimbursement tidak boleh melebihi Rp. 5.000.000")
	}
	errInsert := service.reimbursmentService.Insert(input)
	if errInsert != nil {
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
