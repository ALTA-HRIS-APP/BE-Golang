package service

import (
	"be_golang/klp3/features/target"
	"be_golang/klp3/helper"
	"errors"
	"math"
	"net/url"

	"github.com/go-playground/validator/v10"
)

type targetService struct {
	targetRepo target.TargetDataInterface
	validate   *validator.Validate
}

func New(repo target.TargetDataInterface) target.TargetServiceInterface {
	return &targetService{
		targetRepo: repo,
		validate:   validator.New(),
	}
}

// Create implements target.TargetServiceInterface.
func (s *targetService) Create(input target.TargetEntity) (string, error) {
	userPembuat, err := s.targetRepo.GetUserByIDAPI(input.UserIDPembuat)
	if err != nil {
		return "", err
	}
	userPenerima, err := s.targetRepo.GetUserByIDAPI(input.UserIDPenerima)
	if err != nil {
		return "", err
	}
	err = s.validate.Struct(input)
	if err != nil {
		return "", errors.New("validation error, content target and due date are required")
	}

	if userPembuat.Jabatan == "manager" {
		if userPenerima.Jabatan != "karyawan" {
			return "", errors.New("your role does not have permission to create targets")
		}
		if userPenerima.Devisi != userPembuat.Devisi {
			return "", errors.New("only create targets for same devisi")
		}
	}

	if userPembuat.Jabatan != "c-level" && userPembuat.Jabatan != "manager" {
		return "", errors.New("your role does not have permission to create targets")
	}

	idTarget, err := s.targetRepo.Insert(input)
	if err != nil {
		return "", err
	}
	return idTarget, nil
}

func (s *targetService) GetAll(token string, idUser string, params url.Values) ([]target.TargetEntity, helper.Paginator, error) {

	pagination, err := helper.PaginationBuilder(params.Get("per_page"), params.Get("page"))
	if err != nil {
		return nil, helper.Paginator{}, err
	}
	filter := target.QueryParam{
		Page:         pagination.Page,
		LimitPerPage: pagination.PerPage,
		Offset:       pagination.Offset,
		SearchKonten: params.Get("search_konten"),
	}

	dataTarget, total, err := s.targetRepo.SelectAll(token, filter)
	if err != nil {
		return nil, helper.Paginator{}, err
	}

	paginator := helper.Paginator{
		TotalData: total,
		Page:      pagination.Page,
		TotalPage: int(math.Ceil(float64(total) / float64(pagination.PerPage))),
		PerPage:   pagination.PerPage,
	}
	return dataTarget, paginator, nil

}

// GetById implements target.TargetServiceInterface.
func (s *targetService) GetById(idTarget string, idUser string) (target.TargetEntity, error) {
	result, err := s.targetRepo.Select(idTarget)
	if err != nil {
		return target.TargetEntity{}, err
	}
	return result, nil
}

// UpdateById implements target.TargetServiceInterface.
func (s *targetService) UpdateById(idTarget string, idUser string, targetData target.TargetEntity) error {
	// Get user information who will perform the update
	user, err := s.targetRepo.GetUserByIDAPI(idUser)
	if err != nil {
		return err
	}

	// Get the target to be updated
	existingTarget, err := s.targetRepo.Select(idTarget)
	if err != nil {
		return err
	}

	userTarget, err := s.targetRepo.GetUserByIDAPI(existingTarget.UserIDPenerima)
	if err != nil {
		return err
	}

	allowedToUpdate := false
	if user.Jabatan == "c-level" {
		allowedToUpdate = true
	}

	if user.Jabatan == "manager" {
		if userTarget.Jabatan == "karyawan" || existingTarget.UserIDPenerima == idUser {
			if userTarget.Devisi == user.Devisi {
				allowedToUpdate = true
			}
		}
	}

	if user.Jabatan == "karyawan" {
		if existingTarget.UserIDPenerima == idUser {
			allowedToUpdate = true
		}
	}

	if !allowedToUpdate {
		return errors.New("you do not have permission to edit this target")
	}

	err = s.targetRepo.Update(idTarget, targetData)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById implements target.TargetServiceInterface.
func (s *targetService) DeleteById(idTarget string, idUser string) error {
	user, err := s.targetRepo.GetUserByIDAPI(idUser)
	if err != nil {
		return err
	}

	// Dapatkan target yang akan diperbarui
	existingTarget, err := s.targetRepo.Select(idTarget)
	if err != nil {
		return err
	}

	// Dapatkan pengguna dengan ID sesuai existingTarget.UserIDPenerima
	userTarget, err := s.targetRepo.GetUserByIDAPI(existingTarget.UserIDPenerima)
	if err != nil {
		return err
	}

	allowedToDelete := false

	if user.Jabatan == "c-level" {
		allowedToDelete = true
	}
	if user.Jabatan == "manager" && userTarget.Jabatan == "karyawan" && user.Devisi == userTarget.Devisi {
		allowedToDelete = true
	}

	if !allowedToDelete {
		return errors.New("you do not have permission to delete this target")
	}
	err = s.targetRepo.Delete(idTarget)
	if err != nil {
		return err
	}
	return nil
}
