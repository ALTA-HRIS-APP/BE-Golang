package service

import (
	"be_golang/klp3/features/absensi"
	apinodejs "be_golang/klp3/features/apiNodejs"
	usernodejs "be_golang/klp3/features/userNodejs"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type AbsensiService struct {
	absensiService absensi.AbsensiDataInterface
	validate       *validator.Validate
}

// GetUserByIDAPI implements absensi.AbsensiServiceInterface
func (service *AbsensiService) GetUserByIDAPI(idUser string) (apinodejs.Pengguna, error) {
	// Panggil metode GetUserByIDFromExternalAPI dari lapisan data absensiRepo
	user, err := service.absensiService.GetUserByIDAPI(idUser)
	if err != nil {
		log.Printf("Error consume api in service: %s", err.Error())
		return apinodejs.Pengguna{}, err
	}
	log.Println("consume api in service successfully")
	return user, nil
}

// GetById implements absensi.AbsensiServiceInterface
func (service *AbsensiService) GetById(absensiID string, userID string) (absensi.AbsensiEntity, error) {
	result, err := service.absensiService.SelectById(absensiID, userID)
	if err != nil {
		return absensi.AbsensiEntity{}, err
	}
	return result, nil
}

// Add implements absensi.AbsensiServiceInterface
func (service *AbsensiService) Add(idUser string) error {
	var input absensi.AbsensiEntity

	sekarang := time.Now()

	tanggalWaktuBaru := time.Date(
		sekarang.Year(),
		sekarang.Month(),
		sekarang.Day(),
		sekarang.Hour(),
		sekarang.Minute(),
		sekarang.Second(),
		sekarang.Nanosecond(),
		time.UTC,
	)
	jamSeharusnya := "08:00:00"
	format := "15:04:05"

	t, err := time.Parse(format, jamSeharusnya)
	if err != nil {
		return errors.New("gagal parsing waktu")
	}
	waktuSeharusnyaMasuk := time.Date(sekarang.Year(), sekarang.Month(), sekarang.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
	selisih := tanggalWaktuBaru.Sub(waktuSeharusnyaMasuk)
	keterlambatan := selisih.Minutes()

	keterlambatanInt := int(keterlambatan)
	konvKeterlambatan := strconv.Itoa(keterlambatanInt)

	input.JamMasuk = jamSeharusnya
	input.OverTimeMasuk = konvKeterlambatan
	input.UserID = idUser
	errInsert := service.absensiService.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

// Edit implements absensi.AbsensiServiceInterface
func (service *AbsensiService) Edit(idUser string, id string) error {
	var input absensi.AbsensiEntity
	sekarang := time.Now()

	tanggalWaktuBaru := time.Date(
		sekarang.Year(),
		sekarang.Month(),
		sekarang.Day(),
		sekarang.Hour(),
		sekarang.Minute(),
		sekarang.Second(),
		sekarang.Nanosecond(),
		time.UTC,
	)
	jamSeharusnya := "17:00:00"
	format := "15:04:05"

	t, err := time.Parse(format, jamSeharusnya)
	if err != nil {
		return errors.New("gagal parsing waktu")
	}
	waktuSeharusnyaMasuk := time.Date(sekarang.Year(), sekarang.Month(), sekarang.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
	selisih := tanggalWaktuBaru.Sub(waktuSeharusnyaMasuk)
	keterlambatan := selisih.Minutes()

	keterlambatanInt := int(keterlambatan)
	konvKeterlambatan := strconv.Itoa(keterlambatanInt)

	input.JamKeluar = jamSeharusnya
	input.OverTimePulang = konvKeterlambatan
	errUpdate := service.absensiService.Update(input, idUser, id)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

// Get implements absensi.AbsensiServiceInterface
func (service *AbsensiService) Get(idUser string, param absensi.QueryParams) (bool, []absensi.AbsensiEntity, error) {
	var total_pages int64
	nextPage := true
	dataUser, errUser := usernodejs.GetByIdUser(idUser)
	if errUser != nil {
		return true, nil, errors.New("error get data user")
	}

	if dataUser.Jabatan == "karyawan" {
		count, dataReim, errReim := service.absensiService.SelectAllKaryawan(idUser, param)
		if errReim != nil {
			return true, nil, errReim
		}
		if param.IsClassDashboard {
			total_pages = count / int64(param.ItemsPerPage)
			if count%int64(param.ItemsPerPage) != 0 {
				total_pages += 1
			}

			if param.Page == int(total_pages) {
				nextPage = false
			}
			count_lebih := count / int64(param.Page)
			if count_lebih < int64(param.ItemsPerPage) {
				nextPage = false
			}
		}
		return nextPage, dataReim, nil
	} else {
		count, dataReim, errReim := service.absensiService.SelectAll(param)
		if errReim != nil {
			return true, nil, errReim
		}
		if param.IsClassDashboard {
			total_pages = count / int64(param.ItemsPerPage)
			if count%int64(param.ItemsPerPage) != 0 {
				total_pages += 1
			}

			if param.Page == int(total_pages) {
				nextPage = false
			}
			count_lebih := count / int64(param.Page)
			if count_lebih < int64(param.ItemsPerPage) {
				nextPage = false
			}
		}
		return nextPage, dataReim, nil

	}
}

func New(service absensi.AbsensiDataInterface) absensi.AbsensiServiceInterface {
	return &AbsensiService{
		absensiService: service,
		validate:       validator.New(),
	}
}
