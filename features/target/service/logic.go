package service

import (
	"be_golang/klp3/features/target"
	"errors"

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

	targetID, err := s.targetRepo.Insert(input)
	if err != nil {
		return "", err
	}
	return targetID, nil
}

func (s *targetService) GetAll(userID string, param target.QueryParam) (bool, []target.TargetEntity, error) {
	var totalPage int64
	nextPage := true

	// Get user's role
	user, err := s.targetRepo.GetUserByIDAPI(userID)
	if err != nil {
		return true, nil, err
	}

	var data []target.TargetEntity
	if user.Jabatan == "karyawan" {
		// Karyawan can only view their own targets
		count, karyawanData, err := s.targetRepo.SelectAllKaryawan(userID, param)
		if err != nil {
			return true, nil, err
		}
		if count == 0 {
			nextPage = false
		}
		data = karyawanData
		if param.ExistOtherPage {
			totalPage = count / int64(param.LimitPerPage)
			if count%int64(param.LimitPerPage) != 0 {
				totalPage += 1
			}

			if param.Page == int(totalPage) {
				nextPage = false
			}
			if data == nil {
				nextPage = false
			}
		}
	} else {
		count, allData, err := s.targetRepo.SelectAll(param)
		if err != nil {
			return true, nil, err
		}
		if count == 0 {
			nextPage = false
		}
		data = allData
		if param.ExistOtherPage {
			totalPage = count / int64(param.LimitPerPage)
			if count%int64(param.LimitPerPage) != 0 {
				totalPage += 1
			}

			if param.Page == int(totalPage) {
				nextPage = false
			}
			if data == nil {
				nextPage = false
			}
		}
	}
	return nextPage, data, nil
}

// GetById implements target.TargetServiceInterface.
func (s *targetService) GetById(targetID string, userID string) (target.TargetEntity, error) {
	result, err := s.targetRepo.Select(targetID)
	if err != nil {
		return target.TargetEntity{}, err
	}
	return result, nil
}

// UpdateById implements target.TargetServiceInterface.
func (s *targetService) UpdateById(targetID string, userID string, targetData target.TargetEntity) error {
	// Get user information who will perform the update
	user, err := s.targetRepo.GetUserByIDAPI(userID)
	if err != nil {
		return err
	}

	// Get the target to be updated
	existingTarget, err := s.targetRepo.Select(targetID)
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
		if userTarget.Jabatan == "karyawan" || existingTarget.UserIDPenerima == userID {
			if userTarget.Devisi == user.Devisi {
				allowedToUpdate = true
			}
		}
	}

	if user.Jabatan == "karyawan" {
		if existingTarget.UserIDPenerima == userID {
			allowedToUpdate = true
		}
	}

	if !allowedToUpdate {
		return errors.New("you do not have permission to edit this target")
	}

	err = s.targetRepo.Update(targetID, targetData)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById implements target.TargetServiceInterface.
func (s *targetService) DeleteById(targetID string, userID string) error {
	user, err := s.targetRepo.GetUserByIDAPI(userID)
	if err != nil {
		return err
	}

	// Dapatkan target yang akan diperbarui
	existingTarget, err := s.targetRepo.Select(targetID)
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
	err = s.targetRepo.Delete(targetID)
	if err != nil {
		return err
	}
	return nil
}
