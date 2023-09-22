package service

import (
	"be_golang/klp3/features/cuti"
	usernodejs "be_golang/klp3/features/userNodejs"
	"errors"

	"github.com/go-playground/validator/v10"
)

type CutiService struct {
	cutiService cuti.CutiDataInterface
	validate    *validator.Validate
}

// Delete implements cuti.CutiServiceInterface.
func (service *CutiService) Delete(id string) error {
	err := service.cutiService.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Edit implements cuti.CutiServiceInterface.
func (service *CutiService) Edit(input cuti.CutiEntity, id string, idUser string) error {
	dataUser, errUser := usernodejs.GetByIdUser(idUser)
	if errUser != nil {
		return errors.New("user not found")
	}
	if input.TipeCuti == "melahirkan" {
		if input.JumlahCuti > 90 {
			return errors.New("cuti melahirkan maksimal 90 hari")
		}

	} else if input.TipeCuti == "hari raya" {
		if input.JumlahCuti > 7 {
			return errors.New("cuti hari raya maksimal 7 hari")
		}

	} else {
		if input.JumlahCuti > 12 {
			return errors.New("cuti tahunan maksimal 12 hari")
		}

	}
	dataCuti, errCuti := service.cutiService.SelectById(id)
	if errCuti != nil {
		return errCuti
	}
	dataUserPengaju, errPengaju := usernodejs.GetByIdUser(dataCuti.UserID)
	if errPengaju != nil {
		return errors.New("UserPengaju not found")
	}
	if dataUser.Jabatan == "karyawan" {
		if input.Status != "" {
			return errors.New("karyawan tidak boleh edit status")
		}
		if input.Persetujuan != "" {
			return errors.New("karyawan tidak boleh edit persetujuan")
		}
		input.UserID = idUser
		err := service.cutiService.UpdateKaryawan(input, id)
		if err != nil {
			return err
		}
		return nil
	} else if dataUser.Jabatan == "manager" {
		if input.Status != "" {
			return errors.New("manager tidak boleh edit status")
		}
		if dataUserPengaju.Jabatan == "manager" || dataUserPengaju.Jabatan == "hr" || dataUserPengaju.Jabatan == "c-level" {
			return errors.New("manager hanya berhak approve cuti h-1")
		}
		if input.Persetujuan == "reject" {
			input.Status = "reject"
			err := service.cutiService.Update(input, id)
			if err != nil {
				return err
			}
			return nil
		} else {
			input.Status = "pending (dalam proses)"
			err := service.cutiService.Update(input, id)
			if err != nil {
				return err
			}
			return nil
		}
	} else if dataUser.Jabatan == "hr" {
		if dataUserPengaju.Jabatan == "hr" || dataUserPengaju.Jabatan == "c-level" {
			return errors.New("hr hanya berhak approve cuti h-1")
		}
		if input.Status != "" {
			return errors.New("hr tidak boleh edit status")
		}
		if dataCuti.Status == "pending" {
			return errors.New("cuti belum di approve oleh manager,harap hubungi manager yang bersangkutan terlebih dahulu")
		}
		if input.Persetujuan == "reject" {
			input.Status = "reject"
			err := service.cutiService.Update(input, id)
			if err != nil {
				return err
			}
			return nil
		} else {
			input.Status = "approve"
			err := service.cutiService.Update(input, id)
			if err != nil {
				return err
			}
			return nil
		}
	} else {
		err := service.cutiService.Update(input, id)
		if err != nil {
			return err
		}
		return nil
	}
}

// Get implements cuti.CutiServiceInterface.
func (service *CutiService) Get(token string,idUser string) ([]cuti.CutiEntity, error) {
	dataUser, errUser := usernodejs.GetByIdUser(idUser)
	if errUser != nil {
		return nil, errors.New("eror get data user")
	}
	if dataUser.Jabatan == "karyawan" {
		dataCuti, errCuti := service.cutiService.SelectAllKaryawan(idUser)
		if errCuti != nil {
			return nil, errCuti
		}
		return dataCuti, nil
	} else {
		dataCuti, errCuti := service.cutiService.SelectAll(token)
		if errCuti != nil {
			return nil, errCuti
		}
		return dataCuti, nil
	}
}

// Add implements cuti.CutiServiceInterface.
func (service *CutiService) Add(input cuti.CutiEntity) error {
	const MaxCutiMelahirkan = 90
	const MaxCutiSakit = 3
	const MaxCutiHariRaya = 7
	const MaxCutiTahunan = 12
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validate error")
	}
	if input.TipeCuti == "melahirkan" {
		if input.JumlahCuti > MaxCutiMelahirkan {
			return errors.New("cuti melahirkan maksimal 90 hari")
		}
		err := service.cutiService.Insert(input)
		if err != nil {
			return err
		}
		return nil

	} else if input.TipeCuti == "sakit" {
		if input.JumlahCuti > MaxCutiSakit {
			return errors.New("cuti sakit maksimal 3 hari")
		}
		err := service.cutiService.Insert(input)
		if err != nil {
			return err
		}
		return nil
	} else if input.TipeCuti == "hari raya" {
		if input.JumlahCuti > MaxCutiHariRaya {
			return errors.New("cuti hari raya maksimal 7 hari")
		}
		err := service.cutiService.Insert(input)
		if err != nil {
			return err
		}
		return nil

	} else {
		if input.JumlahCuti > MaxCutiTahunan {
			return errors.New("cuti tahunan maksimal 12 hari")
		}
		err := service.cutiService.Insert(input)
		if err != nil {
			return err
		}
		return nil
	}
}

func New(service cuti.CutiDataInterface) cuti.CutiServiceInterface {
	return &CutiService{
		cutiService: service,
		validate:    validator.New(),
	}
}
