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

func (service *ReimbursementService) EditAdmin(status string, userID string, UserAdmin string, id string) error {
	dataUser, errUser := usernodejs.GetByIdUser(UserAdmin)
	if errUser != nil {
		return errUser
	}
	dataUserA, errUserA := usernodejs.GetByIdUser(userID)
	if errUserA != nil {
		return errUserA
	}

	if dataUser.Jabatan == "manager" {
		if dataUserA.Jabatan == "c-level" {
			return errors.New("tidak dapat approve C-Level, hanya dapat approve di N-1")
		}
		if dataUserA.Jabatan == "HR" {
			return errors.New("tidak dapat approve HR, hanya dapat approve di N-1")
		}
		if dataUserA.Jabatan == "manager" {
			return errors.New("tidak dapat approve manager, hanya dapat approve di N-1")
		}
		errUpdate := service.reimbursmentService.UpdateStatusByManager(status, userID, id)
		if errUpdate != nil {
			return errUpdate
		}

	} else if dataUser.Jabatan == "lead" {
		if dataUserA.Jabatan == "C-Level" {
			return errors.New("tidak dapat approve C-Level, hanya dapat approve di N-1")
		}
		if dataUserA.Jabatan == "HR" {
			return errors.New("tidak dapat approve HR, hanya dapat approve di N-1")
		}
		if dataUserA.Jabatan == "manager" {
			return errors.New("tidak dapat approve manager, hanya dapat approve di N-1")
		}
		if dataUserA.Jabatan == "lead" {
			return errors.New("tidak dapat approve lead, hanya dapat approve di N-1")
		}
		errUpdate := service.reimbursmentService.UpdateStatusByManager(status, userID, id)
		if errUpdate != nil {
			return errUpdate
		}

	} else if dataUser.Jabatan == "HR" {
		if dataUserA.Jabatan == "HR" {
			return errors.New("tidak dapak approve HR, hanya dapat approve di N-1")
		}
		if dataUserA.Jabatan == "C-Level" {
			return errors.New("tidak dapak approve C-Level, hanya dapat approve di N-1")
		}
		errUpdate := service.reimbursmentService.UpdateStatusByHR(status, userID, id)
		if errUpdate != nil {
			return errUpdate
		}
	} else if dataUser.Jabatan == "C-Level" {
		errUpdate := service.reimbursmentService.UpdateStatusByHR(status, userID, id)
		if errUpdate != nil {
			return errUpdate
		}
	} else {
		return errors.New("selain admin dan superadmin, anda tidak dapat merubah status")
	}
	return nil
}

// Edit implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Edit(input reimbusment.ReimbursementEntity, id string) error {
	errUpdate := service.reimbursmentService.UpdateUser(input, input.UserID, id)
	if errUpdate != nil {
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
