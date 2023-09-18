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
func (service *ReimbursementService) Edit(input reimbusment.ReimbursementEntity, id string,idUser string) error {
	dataUser,errUser:=usernodejs.GetByIdUser(idUser)
	if errUser !=nil{
		return errors.New("failed get user by id")
	}

	batasan,errBatasan:=service.reimbursmentService.SelectById(id)
	if errBatasan != nil{
		return errBatasan
	}
	if input.Nominal > batasan{
		return errors.New("nominal tidak boleh melebihi batasan reimbursment")
	}
	if dataUser.Jabatan =="karyawan"{
		if input.BatasanReimburs != 0{
			return errors.New("karyawan tidak berhak mengedit batasan reimbursement, harap berkonsultasi dengan atasan")
		}
		if input.Persetujuan != ""{
			return errors.New("hanya HR yang bisa approve final")
		}
		if input.Status != ""{
			return errors.New("hanya Manager yang bisa approve")
		}
		input.UserID=idUser
		err:=service.reimbursmentService.UpdateKaryawan(input,id)
		if err != nil{
			return err
		}
		return nil
	}else if dataUser.Jabatan == "manager"{
		if input.Persetujuan != ""{
			return errors.New("hanya hr yang bisa approve final")
		}
		err:=service.reimbursmentService.Update(input,id)
		if err != nil{
			return err
		}
		return nil
	}else{
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
