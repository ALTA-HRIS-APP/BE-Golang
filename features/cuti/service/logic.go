package service

import (
	"be_golang/klp3/features/cuti"
	"errors"

	"github.com/go-playground/validator/v10"
)

type CutiService struct {
	cutiService cuti.CutiDataInterface
	validate    *validator.Validate
}

// Add implements cuti.CutiServiceInterface.
func (service *CutiService) Add(input cuti.CutiEntity) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validate error")
	}
	if input.TipeCuti == "melahirkan" {
		if input.JumlahCuti > 90 {
			return errors.New("cuti melahirkan maksimal 90 hari")
		}
		err := service.cutiService.Insert(input)
		if err != nil {
			return err
		}
		return nil

	} else if input.TipeCuti == "hari raya" {
		if input.JumlahCuti > 7 {
			return errors.New("cuti hari raya maksimal 7 hari")
		}
		err := service.cutiService.Insert(input)
		if err != nil {
			return err
		}
		return nil
	} else {
		if input.JumlahCuti > 12 {
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
