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
func (service *ReimbursementService) EditAdmin(status string,userID string,UserAdmin string,id string) error {
	dataUser,errUser:=service.reimbursmentService.SelectUser(UserAdmin)
	if errUser != nil{
		return errUser
	}
	dataUserA,errUserA:=service.reimbursmentService.SelectUser(userID)
	if errUserA != nil{
		return errUserA
	}

	if dataUser.Devisi=="manager"{
		if dataUserA.Devisi =="C-Level"{
			return errors.New("tidak dapat approve C-Level, hanya dapat approve di N-1")
		}
		if dataUserA.Devisi =="HR"{
			return errors.New("tidak dapat approve HR, hanya dapat approve di N-1")
		}
		if dataUserA.Devisi =="manager"{
			return errors.New("tidak dapat approve manager, hanya dapat approve di N-1")
		}
		errUpdate:=service.reimbursmentService.UpdateStatusByManager(status,userID,id)
		if errUpdate != nil{
			return errUpdate
		}

	}else if dataUser.Devisi=="lead"{
		if dataUserA.Devisi =="C-Level"{
			return errors.New("tidak dapat approve C-Level, hanya dapat approve di N-1")
		}
		if dataUserA.Devisi =="HR"{
			return errors.New("tidak dapat approve HR, hanya dapat approve di N-1")
		}
		if dataUserA.Devisi =="manager"{
			return errors.New("tidak dapat approve manager, hanya dapat approve di N-1")
		}
		if dataUserA.Devisi =="lead"{
			return errors.New("tidak dapat approve lead, hanya dapat approve di N-1")
		}		
		errUpdate:=service.reimbursmentService.UpdateStatusByManager(status,userID,id)
		if errUpdate != nil{
			return errUpdate
		}

	}else if dataUser.Devisi=="HR"{
		if dataUserA.Devisi =="HR"{
			return errors.New("tidak dapak approve HR, hanya dapat approve di N-1")
		}
		if dataUserA.Devisi =="C-Level"{
			return errors.New("tidak dapak approve C-Level, hanya dapat approve di N-1")
		}
		errUpdate:=service.reimbursmentService.UpdateStatusByHR(status,userID,id)
		if errUpdate != nil{
			return errUpdate
		}		
	}else if dataUser.Devisi=="C-Level"{
		errUpdate:=service.reimbursmentService.UpdateStatusByHR(status,userID,id)
		if errUpdate != nil{
			return errUpdate
		}	
	}else{
		return errors.New("selain HR dan manager, anda tidak dapat merubah status")
	}
	return nil
}

// Edit implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Edit(input reimbusment.ReimbursementEntity,id string) error {
	errUpdate:=service.reimbursmentService.UpdateUser(input,input.UserID,id)
	if errUpdate != nil{
		return errUpdate
	}

	return nil
}

// Add implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Add(input reimbusment.ReimbursementEntity) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate, data deskripsi, nominal, tipe reimbusment required")
	}
	_, errInsert := service.reimbursmentService.Insert(input)
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
