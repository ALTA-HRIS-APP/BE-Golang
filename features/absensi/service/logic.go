package service

import (
	"be_golang/klp3/features/absensi"
	usernodejs "be_golang/klp3/features/userNodejs"
	"errors"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type AbsensiService struct {
	absensiService absensi.AbsensiDataInterface
	validate       *validator.Validate
}

// SelectById implements absensi.AbsensiServiceInterface
func (service *AbsensiService) SelectById(id string) (absensi.AbsensiEntity, error) {
	return service.absensiService.SelectById(id)
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

// // Add implements absensi.AbsensiServiceInterface
// func (service *AbsensiService) Add(input absensi.AbsensiEntity) error {
// 	errValidate := service.validate.Struct(input)
// 	if errValidate != nil {
// 		return errors.New("validate error")
// 	}

// 	// Tanggal dan waktu awal dalam format time.Time
// 	tanggalWaktuAwal := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC)

// 	// Membuat objek time.Date baru
// 	tanggalWaktuBaru := time.Date(
// 		tanggalWaktuAwal.Year(),
// 		tanggalWaktuAwal.Month(),
// 		tanggalWaktuAwal.Day(),
// 		tanggalWaktuAwal.Hour(),
// 		tanggalWaktuAwal.Minute(),
// 		tanggalWaktuAwal.Second(),
// 		tanggalWaktuAwal.Nanosecond(),
// 		time.UTC,
// 	)

// 	// Menghitung keterlambatan
// 	waktuSeharusnyaMasuk := time.Date(tanggalWaktuAwal.Year(), tanggalWaktuAwal.Month(), tanggalWaktuAwal.Day(), 7, 30, 0, 0, time.UTC)
// 	selisih := tanggalWaktuBaru.Sub(waktuSeharusnyaMasuk)
// 	keterlambatan := selisih.Minutes()

// 	fmt.Printf("Waktu Baru: %s\n", tanggalWaktuBaru)
// 	fmt.Printf("Keterlambatan: %.2f menit\n", keterlambatan)
// 	return nil
// }

func New(service absensi.AbsensiDataInterface) absensi.AbsensiServiceInterface {
	return &AbsensiService{
		absensiService: service,
		validate:       validator.New(),
	}
}
