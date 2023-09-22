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

// GetReimbusherById implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) GetReimbusherById(id string) (reimbusment.ReimbursementEntity, error) {
	data,err:=service.reimbursmentService.SelectById(id)
	if err != nil{
		return reimbusment.ReimbursementEntity{},err
	}
	dataUser,errUser:=service.reimbursmentService.SelectUserById(data.UserID)
	if errUser != nil{
		return reimbusment.ReimbursementEntity{},err
	}
	data.User.ID=dataUser.ID
	data.User.Name=dataUser.NamaLengkap
	
	return data,nil
}

// Delete implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Delete(id string) error {
	err := service.reimbursmentService.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Get implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Get(token string,idUser string, param reimbusment.QueryParams) (bool, []reimbusment.ReimbursementEntity, error) {

	var total_pages int64
	nextPage := true
	dataUser, errUser := service.reimbursmentService.SelectUserById(idUser)
	if errUser != nil {
		return true, nil, errors.New("error get data user")
	}
	if dataUser.Jabatan == "karyawan" {
		count, dataReim, errReim := service.reimbursmentService.SelectAllKaryawan(idUser, param)
		if errReim != nil {
			return true, nil, errReim
		}
		if count == 0 {
			nextPage = false
		}
		if param.IsClassDashboard || count != 0 {
			total_pages = count / int64(param.ItemsPerPage)
			if count%int64(param.ItemsPerPage) != 0 {
				total_pages += 1
			}

			if param.Page == int(total_pages) {
				nextPage = false
			}

			if dataReim == nil {
				nextPage = false
			}
		}
		return nextPage, dataReim, nil
	} else {
		count, dataReim, errReim := service.reimbursmentService.SelectAll(token,param)
		if errReim != nil {
			return true, nil, errReim
		}
		if count == 0 {
			nextPage = false
		}
		if param.IsClassDashboard || count != 0 {
			total_pages = count / int64(param.ItemsPerPage)
			if count%int64(param.ItemsPerPage) != 0 {
				total_pages += 1
			}

			if param.Page == int(total_pages) {
				nextPage = false
			}
			if dataReim == nil {
				nextPage = false
			}
		}
		return nextPage, dataReim, nil

	}
}

// Edit implements reimbusment.ReimbusmentServiceInterface.
func (service *ReimbursementService) Edit(input reimbusment.ReimbursementEntity, id string, idUser string) error {

	dataUser, errUser := service.reimbursmentService.SelectUserById(idUser)
	if errUser != nil {
		return errors.New("error get user")
	}
	dataReimbursement, errBatas := service.reimbursmentService.SelectById(id)
	if errBatas != nil {
		return errBatas
	}
	dataUserPengaju, errUserPengaju := service.reimbursmentService.SelectUserById(dataReimbursement.UserID)
	if errUserPengaju != nil {
		return errors.New("error get user pengaju")
	}
	if dataReimbursement.BatasanReimburs < input.Nominal {
		return errors.New("pengajuan reimbursement tidak boleh melebihi batas")
	}
	if dataUser.Jabatan == "karyawan" {
		if input.Status != "" {
			return errors.New("karyawan tidak boleh mengedit status")
		}
		if input.Persetujuan != "" {
			return errors.New("karyawan tidak boleh mengedit persetujuan")
		}
		if input.BatasanReimburs != 0 {
			return errors.New("karyawan tidak boleh mengedit batas reimbursement")
		}
		input.UserID = idUser
		err := service.reimbursmentService.UpdateKaryawan(input, id)
		if err != nil {
			return err
		}
		return nil
	} else if dataUser.Jabatan == "manager" {
		if dataUser.Devisi != dataUserPengaju.Devisi {
			return errors.New("manager tidak dapat approve karyawan devisi lain")
		}
		if dataUserPengaju.Jabatan == "manager" || dataUserPengaju.Jabatan == "c-level" || dataUserPengaju.Jabatan == "hr" {
			return errors.New("manager hanya bisa approve reimbursement karyawan")
		}
		if input.Status != "" {
			return errors.New("manager tidak boleh mengedit status")
		}
		if input.Persetujuan == "Rejected" {
			input.Status = "Rejected"
			err := service.reimbursmentService.Update(input, id)
			if err != nil {
				return err
			}
			return nil
		} else {
			input.Status = "pending(approve by manager)"
			err := service.reimbursmentService.Update(input, id)
			if err != nil {
				return err
			}
			return nil
		}
	} else if dataUser.Jabatan == "hr" {
		if dataReimbursement.Status == "pending" {
			return errors.New("harus disetujui oleh manager dulu, harap hubungi manager yang bersangkutan")
		}
		if input.Status != "" {
			return errors.New("hr tidak boleh mengedit status")
		}
		if dataUserPengaju.Jabatan == "hr" || dataUserPengaju.Jabatan == "c-level" {
			return errors.New("hanya bisa approve reimbursement karyawan dan manager")
		}
		if input.Persetujuan == "Rejected" {
			input.Status = "Rejected"
			err := service.reimbursmentService.Update(input, id)
			if err != nil {
				return err
			}
			return nil
		} else {
			input.Status = "Done"
			err := service.reimbursmentService.Update(input, id)
			if err != nil {
				return err
			}
			return nil
		}
	} else {
		err := service.reimbursmentService.Update(input, id)
		if err != nil {
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
	if input.Nominal > 5000000 {
		return errors.New("pengajuan reimbursement tidak boleh melebihi Rp. 5.000.000")
	}

	if input.Status != "" {
		return errors.New("tidak dapat menambah status saat create reimbursement")
	}
	if input.Persetujuan != "" {
		return errors.New("tidak dapat menambah persetujuan saat create reimbursement")
	}
	if input.BatasanReimburs != 0 {
		return errors.New("tidak dapat menambah batas reimbursement saat create reimbursement, biarkan default")
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
