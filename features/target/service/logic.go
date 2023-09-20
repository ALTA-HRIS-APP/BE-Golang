package service

import (
	apinodejs "be_golang/klp3/features/apiNodejs"
	"be_golang/klp3/features/target"
	usernodejs "be_golang/klp3/features/userNodejs"

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
	var total_page int64
	var targetID string
	nextPage := true

	// Dapatkan peran pengguna
	user, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		return false, nil, err
	}

	// Dapatkan target yang akan diperbarui
	existingTarget, err := s.targetRepo.Select(targetID, userID)
	if err != nil {
		return false, nil, err
	}

	// Dapatkan pengguna dengan ID sesuai existingTarget.UserIDPenerima
	userTarget, err := usernodejs.GetByIdUser(existingTarget.UserIDPenerima)
	if err != nil {
		return false, nil, err
	}

	// Inisialisasi variabel yang menunjukkan apakah pembaruan diizinkan
	allowedToRead := false
	if user.Jabatan == "c-level" {
		allowedToRead = true
	}

	if user.Jabatan == "manager" {
		if existingTarget.UserIDPenerima == userID {
			allowedToRead = true
		}
		if userTarget.Jabatan == "karyawan" {
			allowedToRead = true
		}
	}
	if user.Jabatan == "karyawan" {
		if existingTarget.UserIDPenerima == userID {
			allowedToRead = true
		}
	}
	// Periksa izin pembaruan
	if !allowedToRead {
		return false, nil, errors.New("anda tidak memiliki izin untuk melihat target ini")
	}

	count, data, err := s.targetRepo.SelectAll(userID, param)
	if err != nil {
		return false, nil, err
	}
	if param.ExistOtherPage {
		total_page = count / int64(param.LimitPerPage)
		if count%int64(param.LimitPerPage) != 0 {
			total_page += 1
		}

		if param.Page == int(total_page) {
			nextPage = false
		}
	}
	return nextPage, data, nil
}

// GetById implements target.TargetServiceInterface.
func (s *targetService) GetById(targetID string, userID string) (target.TargetEntity, error) {
	result, err := s.targetRepo.Select(targetID, userID)
	if err != nil {
		return target.TargetEntity{}, err
	}
	return result, nil
}

// UpdateById implements target.TargetServiceInterface.
func (s *targetService) UpdateById(targetID string, userID string, targetData target.TargetEntity) error {
	// Dapatkan peran pengguna
	user, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		return err
	}

	// Dapatkan target yang akan diperbarui
	existingTarget, err := s.targetRepo.Select(targetID, userID)
	if err != nil {
		return err
	}

	// Dapatkan pengguna dengan ID sesuai existingTarget.UserIDPenerima
	userTarget, err := usernodejs.GetByIdUser(existingTarget.UserIDPenerima)
	if err != nil {
		return err
	}

	// Inisialisasi variabel yang menunjukkan apakah pembaruan diizinkan
	allowedToUpdate := false

	// Pemeriksaan peran pengguna
	if user.Jabatan == "c-level" {
		allowedToUpdate = true
	}

	if user.Jabatan == "manager" {
		// Pemeriksaan apakah manajer dapat mengedit target karyawan atau target milik diri sendiri
		if userTarget.Jabatan == "karyawan" || existingTarget.UserIDPenerima == userID {
			allowedToUpdate = true
		}
	}

	if user.Jabatan == "karyawan" {
		// Pemeriksaan apakah karyawan dapat mengedit target milik diri sendiri
		if existingTarget.UserIDPenerima == userID {
			allowedToUpdate = true
		}
	}

	// Periksa izin pembaruan
	if !allowedToUpdate {
		return errors.New("anda tidak memiliki izin untuk mengedit target ini")
	}

	// Lakukan pembaruan hanya jika diizinkan
	err = s.targetRepo.Update(targetID, userID, targetData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteById implements target.TargetServiceInterface.
func (s *targetService) DeleteById(targetID string, userID string) error {
	// Dapatkan peran pengguna
	user, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		return err
	}

	// Dapatkan target yang akan diperbarui
	existingTarget, err := s.targetRepo.Select(targetID, userID)
	if err != nil {
		return err
	}

	// Dapatkan pengguna dengan ID sesuai existingTarget.UserIDPenerima
	userTarget, err := usernodejs.GetByIdUser(existingTarget.UserIDPenerima)
	if err != nil {
		return err
	}

	// Inisialisasi variabel yang menunjukkan apakah pembaruan diizinkan
	allowedToDelete := false

	if user.Jabatan == "c-level" {
		allowedToDelete = true
	}
	if user.Jabatan == "manager" && userTarget.Jabatan == "karyawan" {
		allowedToDelete = true
	}

	// Periksa izin pembaruan
	if !allowedToDelete {
		return errors.New("anda tidak memiliki izin untuk mengedit target ini")
	}
	err = s.targetRepo.Delete(targetID, userID)
	if err != nil {
		return err
	}
	return nil
}
