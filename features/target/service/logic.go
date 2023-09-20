package service

import (
	apinodejs "be_golang/klp3/features/apiNodejs"
	"be_golang/klp3/features/target"

	"errors"
	"log"

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

func (s *targetService) GetUserByIDAPI(idUser string) (apinodejs.Pengguna, error) {
	// Panggil metode GetUserByIDFromExternalAPI dari lapisan data targetRepo
	user, err := s.targetRepo.GetUserByIDAPI(idUser)
	if err != nil {
		log.Printf("Error consume api in service: %s", err.Error())
		return apinodejs.Pengguna{}, err
	}
	log.Println("consume api in service successfully")
	return user, nil
}

// Create implements target.TargetServiceInterface.
func (s *targetService) Create(input target.TargetEntity) (string, error) {
	userPembuat, err := s.targetRepo.GetUserByIDAPI(input.UserIDPembuat)
	if err != nil {
		log.Printf("Error getting user details for the creator: %s", err.Error())
		return "", err
	}
	userPenerima, err := s.targetRepo.GetUserByIDAPI(input.UserIDPenerima)
	if err != nil {
		log.Printf("Error getting user details for the receiver: %s", err.Error())
		return "", err
	}
	err = s.validate.Struct(input)
	if err != nil {
		log.Printf("Validation error: %s", err.Error())
		return "", errors.New("validation error, content target and due date are required")
	}

	log.Printf("Creator's role: %s", userPembuat.Jabatan)

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
		log.Printf("Error creating target: %s", err.Error())
		return "", err
	}
	log.Println("Target created successfully")
	return targetID, nil
}

func (s *targetService) GetAll(userID string, param target.QueryParam) (bool, []target.TargetEntity, error) {
	var totalPage int64
	var targetID string
	nextPage := true

	// Get user's role
	user, err := s.targetRepo.GetUserByIDAPI(userID)
	if err != nil {
		log.Printf("Error getting user details: %s", err.Error())
		return false, nil, err
	}

	// Get the target to be updated
	existingTarget, err := s.targetRepo.Select(targetID)
	if err != nil {
		log.Printf("Error selecting target: %s", err.Error())
		return false, nil, err
	}

	// Get the user with ID corresponding to existingTarget.UserIDPenerima
	userTarget, err := s.targetRepo.GetUserByIDAPI(existingTarget.UserIDPenerima)
	if err != nil {
		log.Printf("Error getting user details for the target recipient: %s", err.Error())
		return false, nil, err
	}

	// Initialize a variable indicating whether reading is allowed
	allowedToRead := false

	if user.Jabatan == "c-level" {
		allowedToRead = true
	}

	if user.Jabatan == "manager" {
		if existingTarget.UserIDPenerima == userID {
			allowedToRead = true
		}
		if userTarget.Jabatan == "karyawan" && userTarget.Devisi == user.Devisi {
			allowedToRead = true
		}
	}

	if user.Jabatan == "karyawan" {
		if existingTarget.UserIDPenerima == userID {
			allowedToRead = true
		}
	}

	// Check reading permission
	if !allowedToRead {
		log.Println("You do not have permission to view this target.")
		return false, nil, errors.New("you do not have permission to view this target")
	}

	count, data, err := s.targetRepo.SelectAll(userID, param)
	if err != nil {
		log.Printf("Error selecting all targets: %s", err.Error())
		return false, nil, err
	}

	if param.ExistOtherPage {
		totalPage = count / int64(param.LimitPerPage)
		if count%int64(param.LimitPerPage) != 0 {
			totalPage += 1
		}

		if param.Page == int(totalPage) {
			nextPage = false
		}
	}

	log.Println("Targets read successfully")
	return nextPage, data, nil
}

// GetById implements target.TargetServiceInterface.
func (s *targetService) GetById(targetID string, userID string) (target.TargetEntity, error) {
	result, err := s.targetRepo.Select(targetID)
	if err != nil {
		log.Printf("Error selecting target by ID: %s", err.Error())
		return target.TargetEntity{}, err
	}
	log.Println("Target retrieved successfully")
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

	// Get information about the target recipient user
	userTarget, err := s.targetRepo.GetUserByIDAPI(existingTarget.UserIDPenerima)
	if err != nil {
		return err
	}

	// Initialize a variable indicating whether the update is allowed
	allowedToUpdate := false

	// Check permissions based on user role
	if user.Jabatan == "c-level" {
		allowedToUpdate = true
	}

	if user.Jabatan == "manager" {
		if userTarget.Jabatan == "karyawan" || existingTarget.UserIDPenerima == userID {
			// Managers can edit employee targets or their own targets
			// But only if they are in the same division
			if userTarget.Devisi == user.Devisi {
				allowedToUpdate = true
			}
		}
	}

	if user.Jabatan == "karyawan" {
		// Employees can only edit their own targets
		if existingTarget.UserIDPenerima == userID {
			allowedToUpdate = true
		}
	}

	// Check update permission
	if !allowedToUpdate {
		return errors.New("you do not have permission to edit this target")
	}

	// Perform the update only if allowed
	err = s.targetRepo.Update(targetID, targetData)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById implements target.TargetServiceInterface.
func (s *targetService) DeleteById(targetID string, userID string) error {
	//Dapatkan peran pengguna
	user, err := s.targetRepo.GetUserByIDAPI(userID)
	if err != nil {
		log.Printf("Error getting user details: %s", err.Error())
		return err
	}

	// Dapatkan target yang akan diperbarui
	existingTarget, err := s.targetRepo.Select(targetID)
	if err != nil {
		log.Printf("Error selecting target for deletion: %s", err.Error())
		return err
	}

	// Dapatkan pengguna dengan ID sesuai existingTarget.UserIDPenerima
	userTarget, err := s.targetRepo.GetUserByIDAPI(existingTarget.UserIDPenerima)
	if err != nil {
		log.Printf("Error getting user details for the target recipient: %s", err.Error())
		return err
	}

	// Inisialisasi variabel yang menunjukkan apakah pembaruan diizinkan
	allowedToDelete := false

	if user.Jabatan == "c-level" {
		allowedToDelete = true
	}
	if user.Jabatan == "manager" && userTarget.Jabatan == "karyawan" && user.Devisi == userTarget.Devisi {
		allowedToDelete = true
	}

	if !allowedToDelete {
		log.Println("You do not have permission to delete this target.")
		return errors.New("you do not have permission to delete this target")
	}
	err = s.targetRepo.Delete(targetID)
	if err != nil {
		log.Printf("Error deleting target: %s", err.Error())
		return err
	}
	log.Println("Target deleted successfully")
	return nil
}
